package v1

import (
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/account/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// Account routes
func Account(api fiber.Router) {
	// /api/v1/account
	a := api.Group("/account")
	a.Post("/", account.Create)

	a.Post("/salt", limiter.New(limiter.Config{
		Max:               100,              // 100 request / 10 min = 10 rpm
		Expiration:        10 * time.Minute, // 10 min
		LimitReached:      apierr.ErrTooManyRequests,
		Storage:           g.Data,
		LimiterMiddleware: limiter.SlidingWindow{},
	}), account.Salt)

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
