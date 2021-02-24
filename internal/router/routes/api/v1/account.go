package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account"
	"github.com/gofiber/fiber/v2"
)

// Account routes
func Account(api fiber.Router) {
	// /api/v1/account
	a := api.Group("/account")
	// /api/v1/account/auth
	a.Post("/auth", account.Auth)
	// /api/v1/account/logout
	a.Delete("/logout", account.Logout)
}
