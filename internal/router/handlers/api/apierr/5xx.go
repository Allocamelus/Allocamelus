package apierr

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/goutils/fiber/write"
)

// SomthingWentWrong 500
var SomthingWentWrong = New("somthing-went-wrong")

// Err500 middleware
func Err500(c *fiber.Ctx, data interface{}) error {
	return write.JSON(c, 500, data)
}

// ErrSomthingWentWrong 500
func ErrSomthingWentWrong(c *fiber.Ctx) error {
	return Err500(c, SomthingWentWrong)
}
