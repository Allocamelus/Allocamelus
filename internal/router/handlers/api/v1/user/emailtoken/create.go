package emailtoken

import (
	"database/sql"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

var (
	errInvalidEmail = "invalid-email"
)

type createRequest struct {
	Email   string `json:"email" form:"email"`
	Captcha string `json:"captcha" form:"captcha"`
}

type createResp struct {
	Success bool   `json:"success"`
	Error   string `json:"errors,omitempty"`
}

// Create Email Token handler
func Create(c *fiber.Ctx) error {
	request := new(createRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	request.Email = strings.TrimSpace(request.Email)
	if err := user.ValidEmail(request.Email); err != nil {
		return apierr.Err422(c, createResp{Error: errInvalidEmail})
	}

	captchaSolved := true
	if g.Config.HCaptcha.Enabled {
		if err := hcaptcha.Verify(hcaptcha.Values{
			Secret:  g.Data.Config.HCaptcha.Secret,
			Token:   request.Captcha,
			SiteKey: g.Data.Config.HCaptcha.Moderate,
			IP:      clientip.Get(c),
		}); err != nil {
			if err != hcaptcha.ErrInvalidToken {
				logger.Error(err)
				return apierr.ErrSomthingWentWrong(c)
			}
			captchaSolved = false
		}
	}

	if !captchaSolved {
		return apierr.Err401(c, "X-captcha", createResp{Error: "invalid-captcha"})
	}

	userID, err := user.GetIDByEmail(request.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomthingWentWrong(c)
		}
		// Fail silently
		return fiberutil.JSON(c, 200, createResp{Success: true})
	}

	perms, err := user.GetPerms(userID)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}
	if perms != 0 {
		// Fail silently
		return fiberutil.JSON(c, 200, createResp{Success: true})
	}

	tkn, err := token.NewAndInsert(token.Email, userID)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	// Send Email
	go func() {
		logger.Error(tkn.SendEmail(request.Email))
	}()

	return fiberutil.JSON(c, 200, createResp{Success: true})
}