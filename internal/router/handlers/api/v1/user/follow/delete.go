package follow

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// Delete remove follow
func Delete(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Unfollow(session.Context(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}

// Decline follow
func Decline(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Decline(session.Context(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
