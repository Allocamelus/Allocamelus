package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
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
	uniqueName := c.Params("uniqueName")
	if len(uniqueName) == 0 {
		return apierr.ErrNotFound(c)
	}
	userID, err := user.GetIDByUniqueName(uniqueName)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomthingWentWrong(c)
		}
		return apierr.ErrNotFound(c)
	}
	u, err := user.GetPublic(userID)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, getResponse{u})
}
