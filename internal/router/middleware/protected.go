package middleware

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// Protected from non logged in users
func Protected(c *fiber.Ctx) error {
	if !user.LoggedIn(c) {
		return apierr.ErrUnauthorized403(c)
	}
	return c.Next()
}

// ProtectedDecrypter can user session decrypt
func ProtectedDecrypter(c *fiber.Ctx) error {
	if !user.ContextSession(c).CanDecrypt() {
		// TODO: password challenge
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

	following, err := user.Following(user.ContextSession(c).UserID, userID)
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
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return apierr.ErrUnauthorized403(c)
	}

	return checkIdWithSelf(c, ownerId)
}

func checkIdWithSelf(c *fiber.Ctx, userId int64) error {
	if compare.EqualInt64(userId, user.ContextSession(c).UserID) {
		return c.Next()
	}
	return apierr.ErrUnauthorized403(c)
}
