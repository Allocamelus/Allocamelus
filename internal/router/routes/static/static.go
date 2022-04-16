package static

import (
	"path/filepath"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/router/middleware"
	"github.com/gofiber/fiber/v2"
)

var (
	cacheDuration = 24 * time.Hour
	maxAge        = 60 * 60 * 24 * 356 // 1 year
)

var sendIndex = func(c *fiber.Ctx) error {
	return c.SendFile(filepath.Join(g.Config.Path.PublicDir, "index.html"), true)
}

var routes = []string{"/about", "/login", "/signup", "/logout", "/account/verify-email", "/post/*", "/", "/home", "/terms", "/privacy", "/u/*"}

// Static routes
func Static(app *fiber.App) {
	Media(app)
	static := app.Group("/")
	static.Static("/", g.Config.Path.PublicDir, fiber.Static{
		Compress:      true,
		CacheDuration: cacheDuration,
		MaxAge:        maxAge,
	})

	for _, r := range routes {
		static.Get(r, sendIndex)
	}

	static.Get("/*", middleware.NoIndex, func(c *fiber.Ctx) error {
		c.Status(404)
		return c.Next()
	}, sendIndex)
}
