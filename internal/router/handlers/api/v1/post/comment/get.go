package comment

import (
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/pkg/dbutil"
	"github.com/allocamelus/allocamelus/internal/post/comment"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type GetResponse struct {
	Comment *comment.Comment `json:"comment"`
	User    *user.User       `json:"user"`
}

func Get(c *fiber.Ctx) error {
	// Get Comment
	com, errFunc := getCommentForPost(c)
	if errFunc != nil {
		return errFunc(c)
	}

	// Get User
	u, err := user.GetPublic(user.ContextSession(c), com.UserID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	return fiberutil.JSON(c, 200, GetResponse{Comment: com, User: &u})
}

func GetReplies(c *fiber.Ctx) error {
	commentID := fiberutil.ParamsInt64(c, "commentID")
	if commentID == 0 {
		return apierr.ErrNotFound(c)
	}
	page := fiberutil.ParamsInt64(c, "p")
	if page == 0 {
		page = 1
	}

	deep := fiberutil.ParamsBool(c, "deep", defaultDeep)

	// Get Total Replies
	tReplies, err := comment.GetRepliesTotal(commentID, deep)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, tReplies)
	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	replies, err := comment.GetReplies(startNum, perPage, commentID, deep)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	users := new(user.List)
	sessionUser := user.ContextSession(c)
	for _, c := range replies.Comments {
		users.AddUser(sessionUser, c.UserID)
	}

	return fiberutil.JSON(c, 200, GetListResponse{ListComments: replies.ListComments, Users: users.Users, Order: replies.Order})
}

// GetListResponse posts comments
type GetListResponse struct {
	comment.ListComments
	Users user.ListUsers  `json:"users"`
	Order map[int64]int64 `json:"order"`
}

const (
	perPage     int64 = 15
	defaultDeep       = true
)

func GetPostList(c *fiber.Ctx) error {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrNotFound(c)
	}

	page := fiberutil.ParamsInt64(c, "p")
	if page == 0 {
		page = 1
	}

	deep := fiberutil.ParamsBool(c, "deep", defaultDeep)

	// Get Total Comments
	tComments, err := comment.GetPostTotal(postID, deep)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, tComments)

	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	comments, err := comment.GetPostComments(startNum, perPage, postID, deep)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	users := new(user.List)
	sessionUser := user.ContextSession(c)
	for _, c := range comments.Comments {
		users.AddUser(sessionUser, c.UserID)
	}

	return fiberutil.JSON(c, 200, GetListResponse{ListComments: comments.ListComments, Users: users.Users, Order: comments.Order})
}

func getCommentForPost(c *fiber.Ctx) (*comment.Comment, fiber.Handler) {
	// Get post id from params
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return nil, apierr.ErrNotFound
	}
	// Get comment id from params
	commentId := fiberutil.ParamsInt64(c, "commentID")
	if commentId == 0 {
		return nil, apierr.ErrNotFound
	}

	// Get session user from context
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
