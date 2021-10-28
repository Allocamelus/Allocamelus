package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/posts"
	"github.com/allocamelus/allocamelus/internal/router/middleware"
	"github.com/gofiber/fiber/v2"
)

// Posts routes
func Posts(api fiber.Router) {
	// /api/v1/posts
	p := api.Group("/posts", middleware.Protected)
	p.Get("/", posts.Get)
}
