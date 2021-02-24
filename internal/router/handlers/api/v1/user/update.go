package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

// Update user handler
func Update(c *fiber.Ctx) error {
	return write.JSON(c, 200, fiber.Map{"type": "update"})
}
