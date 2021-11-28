package account

import (
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type CreateRequest struct {
	UserName string   `json:"userName" form:"userName"`
	Email    string   `json:"email" form:"email"`
	Password Password `json:"password" form:"password"`
	Key      Key      `json:"key" form:"key"`
	Captcha  string   `json:"captcha" form:"captcha"`
}

type Password struct {
	// Salt must be encoded as $argon2id$v=19$m={int},t={int},p={int}${salt}
	Salt string `json:"salt" form:"salt"`
	Hash string `json:"hash" form:"hash"`
}

type Key struct {
	PublicArmored   string `json:"publicArmored" form:"publicArmored"`
	PrivateArmored  string `json:"privateArmored" form:"privateArmored"`
	RecoveryHash    string `json:"recoveryHash" form:"recoveryHash"`
	RecoveryArmored string `json:"recoveryArmored" form:"recoveryArmored"`
}

func (t *CreateRequest) trimSpace() {
	t.UserName = strings.TrimSpace(t.UserName)
	t.Email = strings.TrimSpace(t.Email)
	t.Password.Salt = strings.TrimSpace(t.Password.Salt)
	t.Password.Hash = strings.TrimSpace(t.Password.Hash)
	t.Captcha = strings.TrimSpace(t.Captcha)
}

type CreateResp struct {
	Success bool        `json:"success"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Create user handler
func Create(c *fiber.Ctx) error {
	if user.LoggedIn(c) {
		return apierr.Err403(c, CreateResp{Errors: []string{"logged-in"}})
	}
	request := new(CreateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.trimSpace()

	newUser := user.New(request.UserName, "", request.Email)

	// Validate newUser
	if userErrs := newUser.ValidatePublic(); !userErrs.Empty() {
		return apierr.Err422(c, CreateResp{Errors: userErrs})
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
			return apierr.Err422(c, CreateResp{Errors: []string{"invalid-captcha"}})
		}
	}

	return fiberutil.JSON(c, 200, request.Password.Hash)
}
