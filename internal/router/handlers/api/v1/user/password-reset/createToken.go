// TODO: reuse less code with email-token

package passwordreset

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
	errInvalidIdentifier = "invalid-identifier"
)

type createRequest struct {
	Identifier string `json:"identifier" form:"identifier"`
	Captcha    string `json:"captcha" form:"captcha"`
}

type createResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// CreateToken Email Token handler
func CreateToken(c *fiber.Ctx) error {
	request := new(createRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	request.Identifier = strings.TrimSpace(request.Identifier)
	if len(request.Identifier) == 0 {
		return err422Invalid(c)
	}

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
			return apierr.Err422(c, createResp{Error: "invalid-captcha"})
		}
	}

	userID, err := user.GetIDByUserName(request.Identifier)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomthingWentWrong(c)
		}
		if err := user.ValidEmail(request.Identifier); err != nil {
			return err422Invalid(c)
		}

		userID, err = user.GetIDByEmail(request.Identifier)
		if err != nil {
			if err != sql.ErrNoRows {
				logger.Error(err)
				return apierr.ErrSomthingWentWrong(c)
			}
			// Fail silently
			return fiberutil.JSON(c, 200, createResp{Success: true})
		}
	}

	// New go routine to create and send email token
	go func() {
		// password Reset token
		tkn, err := token.NewAndInsert(token.Reset, userID)
		if logger.Error(err) {
			return
		}

		// Send Email
		logger.Error(tkn.SendEmail(request.Identifier))
	}()

	return fiberutil.JSON(c, 200, createResp{Success: true})
}

func err422Invalid(c *fiber.Ctx) error {
	return apierr.Err422(c, createResp{Error: errInvalidIdentifier})
}
