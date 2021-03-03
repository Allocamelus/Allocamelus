package user

import (
	"strings"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	userToken "github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type createRequest struct {
	With  string `json:"with" form:"with"`
	Token string `json:"token" form:"token"`
}

type createA10Token struct {
	UniqueName string `json:"uniqueName"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Captcha    string `json:"captcha"`
}

func (t *createA10Token) trimSpace() {
	t.UniqueName = strings.TrimSpace(t.UniqueName)
	t.Name = strings.TrimSpace(t.Name)
	t.Email = strings.TrimSpace(t.Email)
	t.Password = strings.TrimSpace(t.Password)
	t.Captcha = strings.TrimSpace(t.Captcha)
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type createResp struct {
	Success   bool        `json:"success"`
	BackupKey string      `json:"backupKey"`
	Errors    interface{} `json:"errors,omitempty"`
}

const (
	withA10 = "allocamelus"
)

// Create user handler
func Create(c *fiber.Ctx) error {
	if user.LoggedIn(c) {
		return apierr.Err403(c, createResp{Errors: []string{"logged-in"}})
	}
	request := new(createRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	// TODO: add openpgp.js or somthing similar for request encryption w/ client application
	// TODO: add other providers
	if request.With == withA10 {
		var token createA10Token
		if err := json.Unmarshal([]byte(request.Token), &token); err != nil {
			return apierr.Err422(c, apierr.New("invalid-create-token"))
		}
		token.trimSpace()

		newUser := user.New(token.UniqueName, token.Name, token.Email)

		userErrs := make(validation.Errors)
		if err := newUser.ValidatePublic(); err != nil {
			userErrs = err.(validation.Errors)
		}

		userErrs["password"] = newUser.ValidPassword(token.Password)

		if errs := userErrs.Filter(); errs != nil {
			return apierr.Err422(c, createResp{Errors: errs.(validation.Errors)})
		}

		if g.Config.HCaptcha.Enabled {
			if err := hcaptcha.Verify(hcaptcha.Values{
				Secret:  g.Data.Config.HCaptcha.Secret,
				Token:   token.Captcha,
				SiteKey: g.Data.Config.HCaptcha.Moderate,
				IP:      clientip.Get(c),
			}); err != nil {
				if err != hcaptcha.ErrInvalidToken {
					logger.Error(err)
					return apierr.ErrSomthingWentWrong(c)
				}
				return apierr.Err401(c, "X-captcha", createResp{Errors: []string{"invalid-captcha"}})
			}
		}

		if err := newUser.IsEmailUnique(); err != nil {
			time.Sleep(time.Millisecond * (300 + time.Duration(random.FastInt(250))))
			// Fail silently to prevent email leaks
			return fiberutil.JSON(c, 200, createResp{
				Success:   true,
				BackupKey: newUser.BackupKey,
			})
		}

		newUser.GenerateKeys(token.Password)

		backupKey, err := newUser.Insert()
		if logger.Error(err) {
			return apierr.ErrSomthingWentWrong(c)
		}

		emailToken, err := userToken.NewAndInsert(userToken.Email, newUser.ID)
		if logger.Error(err) {
			return apierr.ErrSomthingWentWrong(c)
		}
		// Send Email
		go func() {
			logger.Error(emailToken.SendEmail(newUser.Email))
		}()

		return fiberutil.JSON(c, 200, createResp{Success: true, BackupKey: backupKey})
	}
	return apierr.Err422(c, apierr.New("invalid-with-value"))
}
