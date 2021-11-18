package routes

import (
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/router/routes/api"
	"github.com/allocamelus/allocamelus/internal/router/routes/static"
	"github.com/gofiber/fiber/v2"
)

// Register routes
func Register(app *fiber.App) {
	// /api/v1
	api.V1(app)

	if g.Config.Site.Static {
		// /media & /*
		static.Static(app)
	}
}
