package follow

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// requestsResp
type requestsResp struct {
	Requests map[int64]int64 `json:"requests"`
	Users    user.ListUsers  `json:"users"`
}

// Requests create follow
func Requests(c *fiber.Ctx) error {
	r, err := user.ListRequests(user.ContextSession(c).UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return fiberutil.JSON(c, 200, requestsResp{})
	}

	users := new(user.List)
	sessionUser := user.ContextSession(c)
	for _, userId := range r {
		err = users.AddUser(sessionUser, userId)
		if logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	return fiberutil.JSON(c, 200, requestsResp{Requests: r, Users: users.Users})
}
