package account

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// Auth User authentication handler
func Auth(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, fiber.Map{"type": "auth"})
}
