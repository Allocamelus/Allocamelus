package auth

import (
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// StatusResp struct
type StatusResp struct {
	LoggedIn bool `json:"loggedIn"`
}

// Status handler
func Status(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, StatusResp{LoggedIn: user.LoggedIn(c)})
}
