package update

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type BioRequest struct {
	Bio string `json:"bio" form:"bio"`
}

type BioResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Bio Update handler
func Bio(c *fiber.Ctx) error {
	request := new(BioRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	if err := user.ValidBio(request.Bio); err != nil {
		return apierr.Err422(c, BioResp{Error: err.Error()})
	}

	if err := user.UpdateBio(user.ContextSession(c).UserID, request.Bio); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, BioResp{Success: true})
}
