package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

// Auth User authentication handler
func Auth(c *fiber.Ctx) error {
	return write.JSON(c, 200, fiber.Map{"type": "auth"})
}
