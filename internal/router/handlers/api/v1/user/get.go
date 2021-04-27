package user

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type getResponse struct {
	user.User
}

// Get user handler
func Get(c *fiber.Ctx) error {
	_, userID, hasErr, errApi := shared.GetUserNameIDResp(c)
	if hasErr {
		return errApi
	}

	u, err := user.GetPublic(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, getResponse{u})
}
