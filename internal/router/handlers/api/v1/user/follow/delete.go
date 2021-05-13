package follow

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

type deleteResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Delete remove follow
func Delete(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Unfollow(user.ContextSession(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, deleteResp{Success: true})
}

// Decline follow
func Decline(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Decline(user.ContextSession(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, deleteResp{Success: true})
}
