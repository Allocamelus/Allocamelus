package middleware

import (
	"crypto/subtle"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
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
	}

	currentUserID := user.ContextSession(c).UserID
	if subtle.ConstantTimeEq(int32(userID), int32(currentUserID)) == 0 ||
		subtle.ConstantTimeEq(int32(userID>>32), int32(currentUserID>>32)) == 0 {
		return apierr.ErrUnauthorized403(c)
	}

	return c.Next()
}
