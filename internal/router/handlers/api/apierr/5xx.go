package apierr

import (
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// SomethingWentWrong 500
var SomethingWentWrong = New("something-went-wrong")

// Err500 middleware
func Err500(c *fiber.Ctx, data any) error {
	return fiberutil.JSON(c, 500, data)
}

// ErrSomethingWentWrong 500
func ErrSomethingWentWrong(c *fiber.Ctx) error {
	return Err500(c, SomethingWentWrong)
}
