package emailtoken

import (
	"errors"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

var (
	errInvalidEmail = "invalid-email"
)

type createRequest struct {
	Email   string `json:"email" form:"email"`
	Captcha string `json:"captcha" form:"captcha"`
}

// Create Email Verification Token handler
func Create(c *fiber.Ctx) error {
	request := new(createRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	request.Email = strings.TrimSpace(request.Email)
	// Check for valid request Email
	if err := user.ValidEmail(request.Email); err != nil {
		return apierr.Err422(c, shared.SuccessErrResp{Error: errInvalidEmail})
	}

	// Captcha
	if g.Config.HCaptcha.Enabled {
		if err := hcaptcha.Verify(hcaptcha.Values{
			Secret:  g.Data.Config.HCaptcha.Secret,
			Token:   request.Captcha,
			SiteKey: g.Data.Config.HCaptcha.Moderate,
			IP:      clientip.Get(c),
		}); err != nil {
			if err != hcaptcha.ErrInvalidToken {
				logger.Error(err)
				return apierr.ErrSomethingWentWrong(c)
			}
			return apierr.Err422(c, shared.SuccessErrResp{Error: "invalid-captcha"})
		}
	}

	// Get UserId from database
	userID, err := user.GetIDByEmail(request.Email)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		// Fail silently
		return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
	}

	// Check if user is Verified
	verified, err := user.IsVerified(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	if verified {
		// Fail silently
		return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
	}

	// New go routine to create and send email token
	go func() {
		// Email verification token
		tkn, err := token.NewAndInsert(token.Email, userID)
		if logger.Error(err) {
			return
		}
		// Send Email
		logger.Error(tkn.SendEmail(request.Email))
	}()

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
