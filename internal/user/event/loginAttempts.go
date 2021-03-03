package event

import (
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/gofiber/fiber/v2"
)

// InsertLoginAttempt for user
func InsertLoginAttempt(c *fiber.Ctx, userID int64, pk pgp.PublicKey) error {
	a, err := New(c, LoginAttempt, userID, pk)
	if err != nil {
		return err
	}
	return a.Insert()
}

// GetLoginAttempts for user
func GetLoginAttempts(userID, afterTime int64) (attempts int64, err error) {
	err = preEvents.QueryRow(LoginAttempt, userID, afterTime).Scan(&attempts)
	return
}
