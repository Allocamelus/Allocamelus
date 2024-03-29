package auth

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// StatusResp struct
type StatusResp struct {
	LoggedIn bool      `json:"loggedIn"`
	User     user.User `json:"user,omitempty"`
}

// Status handler
func Status(c *fiber.Ctx) error {
	var resp StatusResp
	resp.LoggedIn = session.LoggedIn(c)
	if resp.LoggedIn {
		var err error
		sessionUser := session.Context(c)
		resp.User, err = user.GetPublic(sessionUser, sessionUser.UserID)
		if logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}
	return fiberutil.JSON(c, 200, resp)
}
