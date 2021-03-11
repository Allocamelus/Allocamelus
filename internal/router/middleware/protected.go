package middleware

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/gofiber/fiber/v2"
)

// Protected from non logged in users
func Protected(c *fiber.Ctx) error {
	if !user.LoggedIn(c) {
		return apierr.ErrUnauthorized403(c)
	}
	return c.Next()
}
