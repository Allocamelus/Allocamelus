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
	_, userID, hasErr, errApi := getUserNameID(c)
	if hasErr {
		return errApi
	}

	u, err := user.GetPublic(userID)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, getResponse{u})
}

func getUserNameID(c *fiber.Ctx) (userName string, userID int64, hasErr bool, errApi error) {
	userName = c.Params("userName")
	if len(userName) == 0 {
		errApi = apierr.ErrNotFound(c)
		hasErr = true
		return
	}
	userID, err := user.GetIDByUserName(userName)
	if err != nil {
		hasErr = true
		if err != sql.ErrNoRows {
			logger.Error(err)
			errApi = apierr.ErrSomthingWentWrong(c)
			return
		}
		errApi = apierr.ErrNotFound(c)
		return
	}
	return
}
