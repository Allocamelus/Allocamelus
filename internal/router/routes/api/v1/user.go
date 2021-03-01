package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user"
	"github.com/gofiber/fiber/v2"
)

// User routes
func User(api fiber.Router) {
	// /api/v1/user
	u := api.Group("/user")
	// /api/v1/user/create
	u.Post("/create", user.Create)
	// /api/v1/user/email-token
	u.Post("/email-token", user.EmailToken)

	// /api/v1/user/:id
	uID := u.Group("/:id")
	uID.Get("/", user.Get)
	// /api/v1/user/:id/delete
	uID.Delete("/delete", user.Delete)
	// /api/v1/user/:id/update
	uID.Delete("/update", user.Update)
}
