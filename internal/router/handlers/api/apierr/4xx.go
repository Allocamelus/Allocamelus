package apierr

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

var (
	// NotFound 404
	NotFound = New("not-found")
)

// Err400 400 Bad Request
func Err400(c *fiber.Ctx, data interface{}) error {
	return write.JSON(c, 400, data)
}

// Err401 401 Unauthorized
func Err401(c *fiber.Ctx, wwwAuth string, data interface{}) error {
	c.Append("WWW-Authenticate", wwwAuth)
	return write.JSON(c, 401, data)
}

// Err404 404 Not Found
func Err404(c *fiber.Ctx, data interface{}) error {
	return write.JSON(c, 404, data)
}

// Err422 Unprocessable Entity
func Err422(c *fiber.Ctx, data interface{}) error {
	return write.JSON(c, 422, data)
}
