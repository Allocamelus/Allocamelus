package auth

import (
	"crypto/subtle"
	"database/sql"
	_ "embed"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/event"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

func init() {
	data.PrepareQueuer.Add(&preGetHash, qGetHash)
}

var (
	//go:embed sql/getHash.sql
	qGetHash   string
	preGetHash *sql.Stmt
)

func getHash(userID int64) (hash string, err error) {
	err = preGetHash.QueryRow(userID).Scan(&hash)
	return
}

var (
	ErrInvalidAuthKey  = errors.New("error user/auth/auth: Invalid Auth Key")
	ErrInvalidUsername = errors.New("error user/auth/login: Invalid Username")
	ErrUnverifiedEmail = errors.New("error user/auth/login: Unverified Email")
)

// CanLogin
// can user (username) login
func CanLogin(username string) (userID int64, err error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return 0, ErrInvalidUsername
	}

	// Check if user exists
	userID, err = user.GetIDByUserName(username)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
		return 0, ErrInvalidUsername
	}

	// Check if user is Verified
	verified, err := user.IsVerified(userID)
	if err != nil {
		return 0, err
	}
	if !verified {
		return 0, ErrUnverifiedEmail
	}
	return userID, nil
}

func AuthKeyLogin(c *fiber.Ctx, userID int64, authKey string) (privateArmored pgp.PrivateKey, err error) {
	if authKey == "" {
		return "", ErrInvalidAuthKey
	}

	// Hash Auth Key
	keyHash, err := HashKey(authKey)
	if err != nil {
		if klog.V(5).Enabled() {
			return
		}
		return "", ErrInvalidAuthKey
	}

	// Get and decode stored Auth Key
	dbHashB64, err := getHash(userID)
	if err != nil {
		return
	}
	dbHash, err := base64.RawStdEncoding.DecodeString(dbHashB64)
	if err != nil {
		return
	}

	// Compare hashes
	if subtle.ConstantTimeCompare(keyHash, dbHash) != 1 {
		return "", ErrInvalidAuthKey
	}

	s, err := session.New(c, userID)
	if err != nil {
		return
	}

	privateArmored, err = key.GetPrivateArmored(userID)
	if err != nil {
		return
	}

	err = s.ToStore(c)
	if err != nil {
		privateArmored = ""
		return
	}

	return
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
