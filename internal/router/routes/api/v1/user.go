package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/emailtoken"
	"github.com/gofiber/fiber/v2"
)

// User routes
func User(api fiber.Router) {
	// /api/v1/user
	u := api.Group("/user")
	u.Post("/", user.Create)

	// /api/v1/user/email-token
	uET := u.Group("/email-token")
	uET.Post("/", emailtoken.Create)
	// /api/v1/user/email-token/validate
	uET.Post("/validate", emailtoken.Validate)

	// /api/v1/user/:id
	uID := u.Group("/:id")
	uID.Get("/", user.Get)
	// /api/v1/user/:id/delete
	uID.Delete("/delete", user.Delete)
	// /api/v1/user/:id/update
	uID.Delete("/update", user.Update)
}
