package static

import (
	"log"

	"github.com/allocamelus/allocamelus/internal/g"
	postMedia "github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/router/middleware"
	"github.com/allocamelus/allocamelus/internal/user/avatar"
	"github.com/gofiber/fiber/v2"
)

// Media routes
func Media(app fiber.Router) {
	media := app.Group("/media", middleware.NoIndex)
	log.Println("logged")
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
