package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user"
	emailtoken "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/email-token"
	passreset "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/password-reset"
	passresetval "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/password-reset/validate"
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

	// /api/v1/user/password-reset
	uPR := u.Group("/password-reset")
	// /api/v1/user/password-reset/token
	uPR.Post("/token", passreset.CreateToken)
	// /api/v1/user/password-reset/validate
	uPRV := uPR.Group("/validate")
	// /api/v1/user/password-reset/validate/token
	uPRV.Post("/token", passresetval.Token)

	// /api/v1/user/:id
	uID := u.Group("/:id")
	uID.Get("/", user.Get)
	// /api/v1/user/:id/delete
	uID.Delete("/delete", user.Delete)
	// /api/v1/user/:id/update
	uID.Delete("/update", user.Update)
}
