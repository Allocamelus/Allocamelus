package comment

import (
	"github.com/allocamelus/allocamelus/internal/post/comment"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type UpdateRequest struct {
	Content string `json:"content" form:"content"`
}

func Update(c *fiber.Ctx) error {
	// Get comment id from params
	commentID := fiberutil.ParamsInt64(c, "commentID")
	if commentID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	// Parse request
	request := new(UpdateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	// Validate Content
	if err := comment.ContentValid(request.Content); err != nil {
		return apierr.Err422(c, shared.SuccessErrResp{Error: err.Error()})
	}

	// Update Content
	if err := comment.UpdateContent(commentID, request.Content); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}

func Delete(c *fiber.Ctx) error {
	// Get comment id from params
	commentID := fiberutil.ParamsInt64(c, "commentID")
	if commentID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	// Delete comment from database
	if err := comment.Delete(commentID); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return c.SendStatus(204)
}
