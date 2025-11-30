package fiberutil

import (
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// JSON is a helper function for writing json
func JSON(c *fiber.Ctx, status int, data any) error {
	if err := c.JSON(data); logger.Error(err) {
		c.Status(500)
		return err
	}
	c.Accepts("application/json")
	return c.SendStatus(status)
}
