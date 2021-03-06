package update

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type NameRequest struct {
	Name string `json:"name" form:"name"`
}

type NameResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Name Update handler
func Name(c *fiber.Ctx) error {
	request := new(NameRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	if err := user.ValidName(request.Name); err != nil {
		return apierr.Err422(c, NameResp{Error: err.Error()})
	}

	if err := user.UpdateName(user.ContextSession(c).UserID, request.Name); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, NameResp{Success: true})
}
