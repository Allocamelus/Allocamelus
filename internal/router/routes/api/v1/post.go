package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/post/update"
	"github.com/allocamelus/allocamelus/internal/router/middleware"
	"github.com/gofiber/fiber/v2"
)

// Post routes
func Post(api fiber.Router) {
	// /api/v1/post
	p := api.Group("/post")
	p.Post("/", middleware.Protected, post.Create)

	// /api/v1/post/:id
	pID := p.Group("/:id")
	pID.Get("/", post.Get)
	// /api/v1/post/:id/publish
	pID.Post("/publish",
		middleware.Protected,
		middleware.ProtectedPosterOnly,
		post.Publish,
	)

	// /api/v1/post/:id/update
	pIDu := pID.Group("/update",
		middleware.Protected,
		middleware.ProtectedPosterOnly,
	)
	// /api/v1/post/:id/update/content
	pIDu.Post("/content", update.Content)
	// /api/v1/post/:id/update/media
	pIDu.Post("/media", update.Media)
}
