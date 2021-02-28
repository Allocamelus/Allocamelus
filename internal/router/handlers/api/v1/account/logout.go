package account

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// Logout handler
func Logout(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, fiber.Map{"type": "logout"})
}
