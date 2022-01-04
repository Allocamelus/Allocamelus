package follow

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// Post create follow
func Post(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Follow(session.Context(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}

// Accept follow
func Accept(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Accept(session.Context(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
