package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

// Logout handler
func Logout(c *fiber.Ctx) error {
	return write.JSON(c, 200, fiber.Map{"type": "logout"})
}
