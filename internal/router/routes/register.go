package routes

import (
	"github.com/allocamelus/allocamelus/internal/router/routes/api"
	"github.com/gofiber/fiber/v2"
)

// Register routes
func Register(app *fiber.App) {
	// /api/v1
	api.V1(app)
}
