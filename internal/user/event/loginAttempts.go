package event

import (
	"context"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
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
	return g.Data.Queries.CountUserEvents(context.Background(), db.CountUserEventsParams{Eventtype: int16(LoginAttempt), Userid: userID, Created: afterTime})
}
