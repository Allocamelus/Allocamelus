package static

import (
	"github.com/allocamelus/allocamelus/internal/g"
	postMedia "github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user/avatar"
	"github.com/gofiber/fiber/v2"
)

// Media routes
func Media(app *fiber.App) {
	media := app.Group("/media")
	media.Static("/"+postMedia.SubPath, g.Config.Path.MediaDir+postMedia.SubPath,
		fiber.Static{
			CacheDuration: cacheDuration,
			MaxAge:        maxAge,
		},
	)

	media.Static("/"+avatar.SubPath, g.Config.Path.MediaDir+avatar.SubPath,
		fiber.Static{
			CacheDuration: cacheDuration,
			MaxAge:        maxAge,
		},
	)
}
