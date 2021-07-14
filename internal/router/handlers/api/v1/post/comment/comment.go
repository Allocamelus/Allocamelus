package comment

import (
	"github.com/allocamelus/allocamelus/internal/pkg/dbutil"
	"github.com/allocamelus/allocamelus/internal/post/comment"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type UpdateRequest struct {
	Content string `json:"content" form:"content"`
}

// PostListResponse posts comments
type PostListResponse struct {
	comment.ListComments
	Users user.ListUsers  `json:"users"`
	Order map[int64]int64 `json:"order"`
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

const perPage int64 = 15

func PostList(c *fiber.Ctx) error {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrNotFound(c)
	}

	page := fiberutil.ParamsInt64(c, "p")
	if page == 0 {
		page = 1
	}

	// Get Total Comments
	tComments, err := comment.GetPostTotal(postID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, tComments)

	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	comments, err := comment.GetPostComments(startNum, perPage, postID, true)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	users := new(user.List)
	sessionUser := user.ContextSession(c)
	for _, c := range comments.Comments {
		users.AddUser(sessionUser, c.UserID)
	}

	return fiberutil.JSON(c, 200, PostListResponse{ListComments: comments.ListComments, Users: users.Users, Order: comments.Order})
}
