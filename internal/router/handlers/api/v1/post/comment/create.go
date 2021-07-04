package comment

import (
	"errors"
	"strconv"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/post/comment"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type CreateRequest struct {
	ReplyTo int64  `json:"replyTo" form:"replyTo"`
	Content string `json:"content" form:"content"`
}

func Create(c *fiber.Ctx) error {
	// Get session user from context
	sUser := user.ContextSession(c)
	if !sUser.Perms.CanPost() {
		return apierr.Err403(c, g.ErrInsufficientPerms.Error())
	}

	// Get post id from params
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	// Parse request
	request := new(CreateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	// Trim Content
	request.Content = strings.TrimSpace(request.Content)
	if err := comment.ContentValid(request.Content); err != nil {
		return apierr.Err422(c, shared.SuccessErrResp{Error: err.Error()})
	}

	// Check if session user can reply to comment
	if err := comment.CanReplyTo(request.ReplyTo, postID, sUser); err != nil {
		switch err {
		// Handle common errors
		case comment.ErrNoComment, post.ErrNoPost, user.ErrViewMeNot:
			if err == comment.ErrNoComment {
				return apierr.Err404(c, shared.SuccessErrResp{Error: "comment-not-found"})
			} else {
				// Log error because middleware should be catching it
				if err == post.ErrNoPost {
					logger.Error(errors.New("api/v1/post/comment/comment: Error post.ErrNoPost ID:" + strconv.Itoa(int(postID)) + " not caught by middleware"))
					return apierr.ErrNotFound(c)
				}

				logger.Error(errors.New("api/v1/post/comment/comment: Error user.ErrViewMeNot ID:" + strconv.Itoa(int(postID)) + " not caught by middleware"))
				return apierr.ErrUnauthorized403(c)
			}
		// Log uncommon errors
		default:
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	// Create comment struct
	comment := comment.New(sUser.UserID, postID, request.ReplyTo, request.Content)
	if err := comment.Validate(); err != nil {
		return apierr.Err422(c, shared.SuccessErrResp{Error: err.Error()})
	}

	// Add comment to database
	if err := comment.Insert(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, shared.SuccessErrResp{Success: true})
}
