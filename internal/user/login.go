package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user/event"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/logger"
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
	privateKey, err := key.GetAndDecryptPK(userID, password)
	if err != nil {
		if err == key.ErrDecryptingKey {
			publickey, err := key.GetPublicKey(userID)
			logger.Error(err)
			event.InsertLoginAttempt(c, userID, publickey)
			return ErrInvalidPassword
		}
		return err
	}

	session, err := NewSession(c, userID, *privateKey)
	if err != nil {
		return err
	}

	return session.ToStore(c)
}

// Logout user logs own errors
func Logout(c *fiber.Ctx) {
	session := g.Session.Get(c)
	session.Delete()
	err := token.DeleteAuth(c)
	if err != sql.ErrNoRows && err != token.ErrAuthCookie {
		logger.Error(err)
	}
}

func authTokenLogin(c *fiber.Ctx) (*Session, error) {
	userToken, err := token.GetAuth(c)
	if err != nil {
		return nil, err
	}

	session, err := NewSessionFromID(c, userToken.UserID)
	if err != nil {
		return nil, err
	}
	if err := session.ToStore(c); err != nil {
		return nil, err
	}
	return session, nil
}
