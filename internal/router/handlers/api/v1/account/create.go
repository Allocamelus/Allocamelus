package account

import (
	"errors"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/auth"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/internal/user/session"
	userToken "github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

type CreateRequest struct {
	UserName string    `json:"userName" form:"userName"`
	Email    string    `json:"email" form:"email"`
	Auth     AuthParts `json:"auth" form:"auth"`
	Key      Key       `json:"key" form:"key"`
	Captcha  string    `json:"captcha" form:"captcha"`
}

type AuthParts struct {
	// Salt must be encoded as $argon2id$v=19$m={int},t={int},p={int}${salt}
	Salt string `json:"salt" form:"salt"`
	Hash string `json:"hash" form:"hash"`
}

type Key struct {
	// PublicArmored armored PGP public key
	PublicArmored string `json:"publicArmored" form:"publicArmored"`
	// PrivateArmored armored PGP private key encrypted with passphrase
	PrivateArmored string `json:"privateArmored" form:"privateArmored"`
	// RecoveryHash hash of recovery key
	RecoveryHash string `json:"recoveryHash" form:"recoveryHash"`
	// RecoveryArmored PGP private key encrypted with recovery key
	RecoveryArmored string `json:"recoveryArmored" form:"recoveryArmored"`
}

type CreateResponse struct {
	Success bool        `json:"success"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Create user handler
func Create(c *fiber.Ctx) error {
	if session.LoggedIn(c) {
		return new(CreateResponse).error(c, []string{"logged-in"})
	}
	request := new(CreateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.trimSpace()

	newUser := user.New(request.UserName, "", request.Email)

	// Validate newUser
	if userErrs := newUser.ValidatePublic(); !userErrs.Empty() {
		return new(CreateResponse).error(c, userErrs.ToString())
	}

	// Validate Password and Key structs
	var passKeyErrs []string
	if err := request.Auth.validate(); err != nil {
		passKeyErrs = append(passKeyErrs, err.Error())
	}
	if err := request.Key.validate(); err != nil {
		passKeyErrs = append(passKeyErrs, err.Error())
	}
	if len(passKeyErrs) > 0 {
		return new(CreateResponse).error(c, passKeyErrs)
	}

	// Check HCaptcha
	if g.Config.HCaptcha.Enabled {
		if err := hcaptcha.Verify(hcaptcha.Values{
			Secret:  g.Data.Config.HCaptcha.Secret,
			Token:   request.Captcha,
			SiteKey: g.Data.Config.HCaptcha.Moderate,
			IP:      clientip.Get(c),
		}); err != nil {
			// Probably not the user's fault
			if err != hcaptcha.ErrInvalidToken {
				logger.Error(err)
				return apierr.ErrSomethingWentWrong(c)
			}
			// Invalid captcha token
			return new(CreateResponse).error(c, []string{"invalid-captcha"})
		}
	}

	var err error

	if err = newUser.IsEmailUnique(); err != nil {
		// Fail silently to prevent email leaks
		return fiberutil.JSON(c, 200, CreateResponse{
			Success: true,
		})
	}

	if err = newUser.Insert(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	keyPrivate := key.NewPrivate()
	keyPrivate.UserID = newUser.ID
	keyPrivate.PublicArmored = pgp.PublicKey(request.Key.PublicArmored)
	keyPrivate.AuthKeyHash, err = auth.HashKeyB64(request.Auth.Hash)
	if err != nil {
		if klog.V(5).Enabled() {
			logger.Error(err)
		}
		return apierr.ErrSomethingWentWrong(c)
	}
	keyPrivate.AuthKeySalt = request.Auth.Salt
	keyPrivate.PrivateArmored = pgp.PrivateKey(request.Key.PrivateArmored)
	keyPrivate.RecoveryKeyHash, err = auth.HashKeyB64(request.Key.RecoveryHash)
	if err != nil {
		if klog.V(5).Enabled() {
			logger.Error(err)
		}
		return apierr.ErrSomethingWentWrong(c)
	}
	keyPrivate.RecoveryArmored = pgp.PrivateKey(request.Key.RecoveryArmored)

	if err = keyPrivate.Insert(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	emailToken, err := userToken.NewAndInsert(userToken.Email, newUser.ID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	// Send Email
	go func() {
		logger.Error(emailToken.SendEmail(newUser.Email))
	}()

	return fiberutil.JSON(c, 200, CreateResponse{Success: true})
}

func (t *CreateRequest) trimSpace() {
	t.UserName = strings.TrimSpace(t.UserName)
	t.Email = strings.TrimSpace(t.Email)
	t.Auth.trimSpace()
	t.Key.trimSpace()
	t.Captcha = strings.TrimSpace(t.Captcha)
}
func (p *AuthParts) trimSpace() {
	p.Salt = strings.TrimSpace(p.Salt)
	p.Hash = strings.TrimSpace(p.Hash)
}
func (k *Key) trimSpace() {
	k.PublicArmored = strings.TrimSpace(k.PublicArmored)
	k.PrivateArmored = strings.TrimSpace(k.PrivateArmored)
	k.RecoveryHash = strings.TrimSpace(k.RecoveryHash)
	k.RecoveryArmored = strings.TrimSpace(k.RecoveryArmored)
}

func (r *CreateResponse) error(c *fiber.Ctx, errs interface{}) error {
	r.Errors = errs
	return apierr.Err422(c, r)
}

var (
	errInvalidPassword = errors.New("invalid-password-struct")
	errInvalidKey      = errors.New("invalid-key-struct")
)

func (p *AuthParts) validate() error {
	if p.Hash == "" || p.Salt == "" {
		return errInvalidPassword
	}
	return nil
}

func (k *Key) validate() error {
	if k.PrivateArmored == "" || k.PublicArmored == "" || k.RecoveryArmored == "" || k.RecoveryHash == "" {
		return errInvalidKey
	}
	return nil
}
