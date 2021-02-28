package user

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// Create user handler
func Create(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, fiber.Map{"type": "create"})
}
