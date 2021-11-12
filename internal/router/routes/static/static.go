package static

import (
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/gofiber/fiber/v2"
)

var (
	cacheDuration = 24 * time.Hour
	maxAge        = 60 * 60 * 24 * 356 // 1 year
)

// Static routes
func Static(app *fiber.App) {
	Media(app)
	static := app.Group("/")
	static.Static("/", g.Config.Path.PublicDir, fiber.Static{
		Compress:      true,
		CacheDuration: cacheDuration,
		MaxAge:        maxAge,
	})
}