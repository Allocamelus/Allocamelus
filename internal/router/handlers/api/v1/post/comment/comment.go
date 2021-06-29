package comment

import (
	"errors"
	"strconv"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/post/comment"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type CreateRequest struct {
	ReplyTo int64  `json:"replyTo" form:"replyTo"`
	Content string `json:"content" form:"content"`
}

type CreateResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type GetResponse struct {
	Comment *comment.Comment `json:"comment"`
	User    *user.User       `json:"user"`
}

func Create(c *fiber.Ctx) error {
	sUser := user.ContextSession(c)
	if !sUser.Perms.CanPost() {
		return apierr.Err403(c, g.ErrInsufficientPerms.Error())
	}

	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	request := new(CreateRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	// Trim Content
	request.Content = strings.TrimSpace(request.Content)
	if err := comment.ContentValid(request.Content); err != nil {
		return apierr.Err422(c, CreateResponse{Error: err.Error()})
	}

	if err := comment.CanReplyTo(request.ReplyTo, sUser); err != nil {
		switch err {
		case comment.ErrNoComment, post.ErrNoPost, user.ErrViewMeNot:
			if err == comment.ErrNoComment {
				return apierr.Err404(c, CreateResponse{Error: "comment-not-found"})
			} else {
				// Log error because middleware should be catching it
				if err == post.ErrNoPost {
					logger.Error(errors.New("api/v1/post/comment/comment: Error post.ErrNoPost ID:" + strconv.Itoa(int(postID)) + " not caught by middleware"))
					return apierr.ErrNotFound(c)
				}

				logger.Error(errors.New("api/v1/post/comment/comment: Error user.ErrViewMeNot ID:" + strconv.Itoa(int(postID)) + " not caught by middleware"))
				return apierr.ErrUnauthorized403(c)
			}
		default:
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	comment := comment.New(sUser.UserID, postID, request.ReplyTo, request.Content)
	if err := comment.Validate(); err != nil {
		return apierr.Err422(c, CreateResponse{Error: err.Error()})
	}

	if err := comment.Insert(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, CreateResponse{Success: true})
}

func Get(c *fiber.Ctx) error {
	com, errFunc := getCommentForPost(c)
	if errFunc != nil {
		return errFunc(c)
	}

	u, err := user.GetPublic(user.ContextSession(c), com.UserID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	return fiberutil.JSON(c, 200, GetResponse{Comment: com, User: &u})
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}

func GetReplies(c *fiber.Ctx) error {
	return nil
}

func PostList(c *fiber.Ctx) error {
	return nil
}

func getCommentForPost(c *fiber.Ctx) (*comment.Comment, fiber.Handler) {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return nil, apierr.ErrNotFound
	}
	commentId := fiberutil.ParamsInt64(c, "commentID")
	if commentId == 0 {
		return nil, apierr.ErrNotFound
	}

	s := user.ContextSession(c)
	com, err := comment.GetForUser(commentId, s)
	if err != nil {
		if err != comment.ErrNoComment {
			if err == user.ErrViewMeNot {
				return nil, apierr.ErrUnauthorized403
			}
			logger.Error(err)
			return nil, apierr.ErrSomethingWentWrong
		}
		return nil, apierr.ErrNotFound
	}

	// Allow user to view own comments
	if !compare.EqualInt64(com.UserID, s.UserID) {
		// Otherwise only allow comments to be viewed with their post
		if !compare.EqualInt64(com.PostID, postID) {
			return nil, apierr.ErrUnauthorized403
		}
	}
	return com, nil
}
