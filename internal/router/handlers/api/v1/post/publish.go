package post

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// Publish post handler
func Publish(c *fiber.Ctx) error {
	p, errFunc := getForUser(c)
	if errFunc != nil {
		return errFunc(c)
	}

	if err := p.Publish(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
