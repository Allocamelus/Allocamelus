package middleware

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/post/comment"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// Protected from non logged in users
func Protected(c *fiber.Ctx) error {
	if !session.LoggedIn(c) {
		return apierr.ErrUnauthorized403(c)
	}
	return c.Next()
}

// ProtectedSelfOnly only allow access to user with same username as :userName
//
// Checks id's of :userName & user session
func ProtectedSelfOnly(c *fiber.Ctx) error {
	_, userID, errApi := shared.GetUserNameAndID(c)
	if errApi == apierr.SomethingWentWrong {
		return apierr.ErrSomethingWentWrong(c)
	} else if errApi == apierr.NotFound {
		return apierr.ErrUnauthorized403(c)
	}

	return checkIdWithSelf(c, userID)
}

// ProtectedPubOrFollow only allow access to user if public or is being followed
// or is self
//
// Checks id's of :userName & user session
func ProtectedPubOrFollow(c *fiber.Ctx) error {
	_, userID, errApi := shared.GetUserNameAndID(c)
	if errApi == apierr.SomethingWentWrong {
		return apierr.ErrSomethingWentWrong(c)
	} else if errApi == apierr.NotFound {
		return apierr.ErrUnauthorized403(c)
	}

	userType, err := user.GetType(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	if userType.Public() {
		return c.Next()
	}

	following, err := user.Following(session.Context(c).UserID, userID)
	if logger.Error(err) {
		// return if following silently
		return apierr.ErrSomethingWentWrong(c)
	}
	if !following.Following {
		return apierr.ErrUnauthorized403(c)
	}

	return c.Next()
}

// ProtectedPosterOnly only allow access to post owner
func ProtectedPosterOnly(c *fiber.Ctx) error {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	ownerId, err := post.GetUserId(postID)
	return sessionIdCheck(c, ownerId, err)
}

// ProtectedCanViewPost only allow access to user who can view post
func ProtectedCanViewPost(c *fiber.Ctx) error {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	err := post.CanView(postID, session.Context(c))
	if err != nil {
		if err == post.ErrNoPost {
			return apierr.ErrNotFound(c)
		}
		if err == user.ErrViewMeNot {
			return apierr.ErrUnauthorized403(c)
		}
		logger.Error(err)
		return apierr.ErrSomethingWentWrong(c)
	}

	return c.Next()
}

// ProtectedCommentPost only allow comment to be viewed with
func ProtectedCommentPost(c *fiber.Ctx) error {
	// Get url post id
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return apierr.ErrUnauthorized403(c)
	}
	// Get url comment id
	commentID := fiberutil.ParamsInt64(c, "commentID")
	if commentID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	// Get comment post id
	cPostID, err := comment.GetPostID(commentID)
	if err != nil {
		// If No comment
		if err == comment.ErrNoComment {
			return apierr.ErrNotFound(c)
		}

		logger.Error(err)
		return apierr.ErrSomethingWentWrong(c)
	}

	// Check url id with comment post id
	if compare.EqualInt64(postID, cPostID) {
		return c.Next()
	}

	return apierr.ErrUnauthorized403(c)
}

func ProtectedCommenterOnly(c *fiber.Ctx) error {
	commentID := fiberutil.ParamsInt64(c, "commentID")
	if commentID == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	ownerId, err := comment.GetUserId(commentID)
	return sessionIdCheck(c, ownerId, err)
}

func sessionIdCheck(c *fiber.Ctx, userId int64, err error) error {
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return apierr.ErrUnauthorized403(c)
	}

	return checkIdWithSelf(c, userId)
}

func checkIdWithSelf(c *fiber.Ctx, userId int64) error {
	if compare.EqualInt64(userId, session.Context(c).UserID) {
		return c.Next()
	}
	return apierr.ErrUnauthorized403(c)
}
