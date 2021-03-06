package emailtoken

import (
	"strings"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

var (
	errInvalidToken = "invalid-token"
	errExpiredToken = "expired-token"
)

type validateRequest struct {
	Selector string `json:"selector" form:"selector"`
	Token    string `json:"token" form:"token"`
}

type validateResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Validate Email Token handler
func Validate(c *fiber.Ctx) error {
	request := new(validateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.Selector = strings.TrimSpace(request.Selector)
	request.Token = strings.TrimSpace(request.Token)

	// Check for email token
	tkn, err := token.Check(request.Selector, request.Token, token.Email)
	if err != nil {
		if err == token.ErrExpiredToken {
			return apierr.Err422(c, validateResp{Error: errExpiredToken})
		}
		return apierr.Err422(c, validateResp{Error: errInvalidToken})
	}

	// Update type from Unverified to Private
	if err := user.UpdateType(tkn.UserID, user.Private); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	// Remove Token
	if err := tkn.Delete(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	return fiberutil.JSON(c, 200, validateResp{Success: true})
}
