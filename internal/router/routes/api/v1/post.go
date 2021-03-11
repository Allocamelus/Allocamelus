package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/post"
	"github.com/gofiber/fiber/v2"
)

// Post routes
func Post(api fiber.Router) {
	// /api/v1/post
	p := api.Group("/post")
	p.Post("/", post.Create)

	pID := p.Group("/:id")
	pID.Get("/", post.Get)
}
