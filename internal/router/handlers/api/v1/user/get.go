package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

// Get user handler
func Get(c *fiber.Ctx) error {
	return write.JSON(c, 200, fiber.Map{"type": "get"})
}
