package update

import (
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type ContentRequest struct {
	Content string `json:"content" form:"content"`
}

// Content Update handler
func Content(c *fiber.Ctx) error {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	request := new(ContentRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	if err := post.ValidateContent(request.Content); err != nil {
		return apierr.Err422(c, shared.SuccessErrResp{Error: err.Error()})
	}

	if err := post.UpdateContent(postID, request.Content); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
