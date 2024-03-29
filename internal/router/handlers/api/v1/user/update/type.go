package update

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type TypeRequest struct {
	Type user.Types `json:"type" form:"type"`
}

// Type Update handler
func Type(c *fiber.Ctx) error {
	request := new(TypeRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	var newType user.Types
	switch request.Type {
	case user.Public:
		newType = user.Public
	default:
		newType = user.Private
	}

	ctxUser := session.Context(c)
	if err := user.UpdateType(ctxUser.UserID, newType); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	if newType.Public() {
		if err := user.AcceptAll(ctxUser.UserID); logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
