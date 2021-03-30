package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/meta/captcha"
	"github.com/gofiber/fiber/v2"
)

func Meta(api fiber.Router) {
	m := api.Group("/meta")
	m.Get("/captcha/site-keys", captcha.SiteKeys)
}
