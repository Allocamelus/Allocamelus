package user

import (
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/user/event"
	"github.com/gofiber/fiber/v2"
)

// LoginDifficulty int8
type LoginDifficulty int8

const (
	// None user has 0-2 login attempts
	None LoginDifficulty = iota
	// Easy user has 3-4 login attempts
	Easy
	// Medium user has 5-6 login attempts
	Medium
	// Hard user has 7-8 login attempts
	Hard
	// ExtraHard user has 9+ login attempts
	ExtraHard
)

var attemptCountDuration = time.Hour * 2

// LoginDiff user's login difficulty
// based on login attempts
func LoginDiff(userID int64) (LoginDifficulty, error) {
	afterTime := time.Now().Unix() - int64(attemptCountDuration.Seconds())
	attempts, err := event.GetLoginAttempts(userID, afterTime)
	if err != nil {
		return None, err
	}
	if attempts <= 2 {
		return None, nil
	} else if attempts <= 4 {
		return Easy, nil
	} else if attempts <= 6 {
		return Medium, nil
	} else if attempts <= 8 {
		return Hard, nil
	}
	return ExtraHard, nil
}

// ErrInvalidPassword Invalid User Password
var ErrInvalidPassword = errors.New("user/login: Invalid User Password")

// PasswordLogin Attempt
func PasswordLogin(c *fiber.Ctx, userID int64, password string) error {
	privateKey, err := GetAndDecryptPK(userID, password)
	if err != nil {
		if err == ErrDecryptingKey {
			return ErrInvalidPassword
		}
		return err
	}
	session, err := NewSessionFromID(c, userID, *privateKey)
	if err != nil {
		return err
	}
	session.ToStore(c)
	return nil
}
