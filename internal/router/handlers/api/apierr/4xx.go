package apierr

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

var (
	// InvalidRequestParams 400
	InvalidRequestParams = New("invalid-request-parameters")
	// NotFound 404
	NotFound = New("not-found")
	// Unauthorized403 because not using www-Authenticate headers
	Unauthorized403 = New("unauthorized-403")
	// UnprocessableEntity 422
	UnprocessableEntity = New("unprocessable-entity")
	// TooManyRequests 429
	TooManyRequests = New("too-many-requests")
)

// Err400 400 Bad Request
func Err400(c *fiber.Ctx, data interface{}) error {
	return fiberutil.JSON(c, 400, data)
}

// ErrInvalidRequestParams 400
func ErrInvalidRequestParams(c *fiber.Ctx) error {
	return Err400(c, InvalidRequestParams)
}

// Err401 401 Unauthorized
func Err401(c *fiber.Ctx, wwwAuth string, data interface{}) error {
	c.Append("WWW-Authenticate", wwwAuth)
	return fiberutil.JSON(c, 401, data)
}

// Err403 403 Forbidden
func Err403(c *fiber.Ctx, data interface{}) error {
	return fiberutil.JSON(c, 403, data)
}

// ErrUnauthorized403 403
func ErrUnauthorized403(c *fiber.Ctx) error {
	return Err403(c, Unauthorized403)
}

// Err404 404 Not Found
func Err404(c *fiber.Ctx, data interface{}) error {
	return fiberutil.JSON(c, 404, data)
}

// ErrNotFound 404
func ErrNotFound(c *fiber.Ctx) error {
	return Err404(c, NotFound)
}

// Err422 Unprocessable Entity
func Err422(c *fiber.Ctx, data interface{}) error {
	return fiberutil.JSON(c, 422, data)
}

// ErrUnprocessableEntity 422
func ErrUnprocessableEntity(c *fiber.Ctx) error {
	return Err422(c, UnprocessableEntity)
}

// Err429 Too Many Requests
func Err429(c *fiber.Ctx, data interface{}) error {
	return fiberutil.JSON(c, 429, data)
}

// ErrTooManyRequests 429
func ErrTooManyRequests(c *fiber.Ctx) error {
	return Err429(c, TooManyRequests)
}
