package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account/auth"
	"github.com/gofiber/fiber/v2"
)

// Account routes
func Account(api fiber.Router) {
	// /api/v1/account
	a := api.Group("/account")
	a.Post("/", account.Create)

	a.Post("/salt", account.Salt)

	// /api/v1/account/auth
	aA := a.Group("/auth")
	aA.Post("/", account.Auth)
	// /api/v1/account/auth/keep-alive
	aA.Post("/keep-alive", auth.KeepAlive)
	// /api/v1/account/auth/status
	aA.Get("/status", auth.Status)

	// /api/v1/account/logout
	a.Delete("/logout", account.Logout)
}
