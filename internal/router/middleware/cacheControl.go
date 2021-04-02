package middleware

import (
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/gofiber/fiber/v2"
)

// CacheControl do caches
func CacheControl(c *fiber.Ctx) error {
	if user.LoggedIn(c) {
		c.Append("cache-control", "private")
	} else {
		c.Append("cache-control", "no-store, no-cache, must-revalidate")
	}
	return c.Next()
}
