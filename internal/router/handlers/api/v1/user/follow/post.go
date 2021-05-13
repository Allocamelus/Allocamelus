package follow

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

type postResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Post create follow
func Post(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Follow(user.ContextSession(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, postResp{Success: true})
}

// Accept follow
func Accept(c *fiber.Ctx) error {
	_, userID, hasErr, err := shared.GetUserNameIDResp(c)
	if hasErr {
		return err
	}

	err = user.Accept(user.ContextSession(c).UserID, userID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, postResp{Success: true})
}
