package middleware

import (
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// Session saving middelware
func Session(c *fiber.Ctx) error {
	user.SessionToContext(c)
	// next routes
	if err := c.Next(); !logger.Error(err) {
		g.Session.Set(c)
	}
	return nil
}
