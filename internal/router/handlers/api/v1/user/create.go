package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

// Create user handler
func Create(c *fiber.Ctx) error {
	return write.JSON(c, 200, fiber.Map{"type": "create"})
}
