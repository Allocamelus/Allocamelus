package shared

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// GetUserNameAndID from :userName
func GetUserNameAndID(c *fiber.Ctx) (userName string, userID int64, errApi apierr.APIErr) {
	userName = c.Params("userName")
	if userName == "" {
		errApi = apierr.NotFound
		return
	}
	userID, err := user.GetIDByUserName(userName)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			errApi = apierr.SomethingWentWrong
			return
		}
		errApi = apierr.NotFound
		return
	}
	return
}

func GetUserNameIDResp(c *fiber.Ctx) (string, int64, bool, error) {
	userName, userId, errApi := GetUserNameAndID(c)
	if !errApi.Empty() {
		switch errApi {
		case apierr.NotFound:
			return "", 0, true, apierr.ErrNotFound(c)
		default:
			return "", 0, true, apierr.ErrSomethingWentWrong(c)
		}
	}
	return userName, userId, false, nil
}
