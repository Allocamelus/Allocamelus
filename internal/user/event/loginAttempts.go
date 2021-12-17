package event

import (
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/gofiber/fiber/v2"
)

// InsertLoginAttempt for user
func InsertLoginAttempt(c *fiber.Ctx, userID int64, pk ...*key.Public) error {
	_, err := InsertNew(c, LoginAttempt, userID, pk...)
	return err
}

// GetLoginAttempts for user
func GetLoginAttempts(userID, afterTime int64) (attempts int64, err error) {
	err = preEvents.QueryRow(LoginAttempt, userID, afterTime).Scan(&attempts)
	return
}
