package middleware

import "github.com/gofiber/fiber/v2"

// NoIndex x-robots-tag
func NoIndex(c *fiber.Ctx) error {
	c.Append("x-robots-tag", "noindex")
	return c.Next()
}
