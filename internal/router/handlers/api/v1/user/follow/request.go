package follow

import (
	"errors"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// requestsResp
type requestsResp struct {
	Requests map[int64]int64 `json:"requests"`
	Users    user.ListUsers  `json:"users"`
}

// Requests create follow
func Requests(c *fiber.Ctx) error {
	r, err := user.ListRequests(session.Context(c).UserID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return fiberutil.JSON(c, 200, requestsResp{})
	}

	users := new(user.List)
	sessionUser := session.Context(c)
	for _, userId := range r {
		err = users.AddUser(sessionUser, userId)
		if logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	return fiberutil.JSON(c, 200, requestsResp{Requests: r, Users: users.Users})
}
