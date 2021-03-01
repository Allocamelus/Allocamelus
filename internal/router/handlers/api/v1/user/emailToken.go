package user

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

func emailToken(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, fiber.Map{"type": "email-verify"})
}
