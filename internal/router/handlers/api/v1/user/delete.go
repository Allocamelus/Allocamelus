package user

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// Delete User handler
// TODO
func Delete(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, fiber.Map{"type": "delete"})
}
