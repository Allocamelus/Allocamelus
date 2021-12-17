package account

import (
	"github.com/allocamelus/allocamelus/internal/user/auth"
	"github.com/gofiber/fiber/v2"
)

// Logout handler
func Logout(c *fiber.Ctx) error {
	auth.Logout(c)
	return c.SendStatus(204)
}
