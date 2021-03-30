package api

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	v1 "github.com/allocamelus/allocamelus/internal/router/routes/api/v1"
	"github.com/gofiber/fiber/v2"
)

// V1 api routes
func V1(app *fiber.App) {
	// /api/v1
	api := app.Group("/api/v1")

	// /api/v1/account
	v1.Account(api)

	// /api/v1/meta
	v1.Meta(api)

	// /api/v1/user
	v1.User(api)

	// /api/v1/post
	v1.Post(api)

	// /api/v1/posts
	v1.Posts(api)

	// 404 Error
	api.Use(apierr.ErrNotFound)
}
