package middleware

import (
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/logger"
)

// Session saving middelware
func Session(c *fiber.Ctx) error {
	// next routes
	if err := c.Next(); !logger.Error(err) {
		g.Session.Set(c)
	}
	return nil
}
