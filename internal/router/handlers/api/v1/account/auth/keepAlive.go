package auth

import (
	"github.com/gofiber/fiber/v2"
)

// KeepAlive handler
func KeepAlive(c *fiber.Ctx) error {
	return c.SendStatus(204)
}
