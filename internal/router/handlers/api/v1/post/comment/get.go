package comment

import (
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
	com, errFunc := getComment(c)
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

	// Get depth from query
	depth := fiberutil.ParamsInt64(c, "depth", defaultDepth)

	// Get Total Replies
	tReplies, err := comment.GetRepliesTotal(commentID, depth)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(topPerPage, page, tReplies)
	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	replies, err := comment.GetReplies(startNum, topPerPage, commentID, depth)
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
	topPerPage   int64 = 10
	maxPerPage   int64 = 40
	defaultDepth int64 = 3
)

func GetPostList(c *fiber.Ctx) error {
	// Get id (postID) from query
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrNotFound(c)
	}

	// Get p (page) from query
	page := fiberutil.ParamsInt64(c, "p")
	if page == 0 {
		page = 1
	}

	// Get depth from query
	depth := fiberutil.ParamsInt64(c, "depth", defaultDepth)

	// Get Total top level Comments
	tComments, err := comment.GetPostTopLevel(postID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	// Page calculations are done for top level comments only
	startNum, totalPages := dbutil.ItemPageCalc(topPerPage, page, tComments)
	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	// Get comments
	comments, err := comment.GetPostComments(startNum, topPerPage, maxPerPage, postID, depth)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	users := new(user.List)
	sessionUser := user.ContextSession(c)
	// Get users for comments
	for _, c := range comments.Comments {
		users.AddUser(sessionUser, c.UserID)
	}

	return fiberutil.JSON(c, 200, GetListResponse{ListComments: comments.ListComments, Users: users.Users, Order: comments.Order})
}

type GetTotal struct {
	Total int64 `json:"total"`
}

func GetTotalForPost(c *fiber.Ctx) error {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrNotFound(c)
	}

	// Get Total Comments
	tComments, err := comment.GetPostTotal(postID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	return fiberutil.JSON(c, 200, GetTotal{Total: tComments})
}

func getComment(c *fiber.Ctx) (*comment.Comment, fiber.Handler) {
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

	return com, nil
}
