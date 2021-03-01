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
	UserID  int64  `json:"userId,omitempty"`
	Error   string `json:"errors,omitempty"`
}

// Validate Email Token handler
func Validate(c *fiber.Ctx) error {
	request := new(validateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.Selector = strings.TrimSpace(request.Selector)
	request.Token = strings.TrimSpace(request.Token)
	tkn, err := token.Check(request.Selector, request.Token, token.Email)
	if err != nil {
		var respErr string
		if err == token.ErrExpiredToken {
			respErr = errExpiredToken
		} else {
			respErr = errInvalidToken
		}
		return apierr.Err422(c, validateResp{Error: respErr})
	}
	if err := user.UpdatePerms(tkn.UserID, user.DefaultPerms); logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}
	userID := tkn.UserID
	if err := tkn.Delete(); logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}
	return fiberutil.JSON(c, 200, validateResp{Success: true, UserID: userID})
}
