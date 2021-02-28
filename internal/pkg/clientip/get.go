package clientip

import (
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/gofiber/fiber/v2"
)

// Get client ip from context
func Get(c *fiber.Ctx) string {
	if g.Config.Cloudflare.Enabled {
		if len(c.Get("CF-Connecting-IP")) > 0 {
			return c.Get("CF-Connecting-IP")
		}
	}

	if len(c.IPs()) >= 1 {
		return c.IPs()[0]
	}

	return c.IP()
}
