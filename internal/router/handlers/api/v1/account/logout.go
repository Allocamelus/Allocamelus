package account

import (
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/gofiber/fiber/v2"
)

// Logout handler
func Logout(c *fiber.Ctx) error {
	user.Logout(c)
	return c.SendStatus(204)
}
