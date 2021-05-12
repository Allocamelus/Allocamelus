package user

import (
	"bytes"
	"crypto/sha512"
	"crypto/subtle"
	"database/sql"
	"errors"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

const storeName = "session"

// NewSession new session
func NewSession(c *fiber.Ctx, userID int64, privateKey pgp.PrivateKey) (*Session, error) {
	var err error
	session := new(Session)
	session.LoggedIn = true
	session.UserID = userID
	session.UserName, err = GetUserNameByID(userID)
	if err != nil {
		return nil, err
	}

	session.PrivateKey = privateKey
	pkSalt, err := key.GetPrivateKeySalt(userID)
	if err != nil {
		return nil, err
	}

	perms, err := GetPerms(userID)
	if err != nil {
		return nil, err
	}

	session.Perms = perms
	session.LoginToken = genLoginToken(c, []byte(pkSalt))
	return session, nil
}

// NewSessionFromID new session from user id
// Can not Decrypt
func NewSessionFromID(c *fiber.Ctx, userID int64) (*Session, error) {
	return NewSession(c, userID, pgp.PrivateKey{})
}

// SessionToContext set user session to context
func SessionToContext(c *fiber.Ctx) {
	c.Locals(storeName, GetSession(c))
}

// ContextSession get user session from fiber context
func ContextSession(c *fiber.Ctx) *Session {
	session := c.Locals(storeName)
	if session != nil {
		return session.(*Session)
	}
	return &Session{}
}

// LoggedIn User
func LoggedIn(c *fiber.Ctx) bool {
	return ContextSession(c).LoggedIn
}

// ToStore set user session to session store
func (s *Session) ToStore(c *fiber.Ctx) error {
	if s != nil {
		sessionBytes, err := s.MarshalMsg(nil)
		if err != nil {
			return err
		}

		store := g.Session.Get(c)
		store.Regenerate()
		return store.Set(storeName, sessionBytes)
	}
	return errors.New("session/session: nil *Session")
}

// GetSession get session from store
// attempt auth login on New session
func GetSession(c *fiber.Ctx) *Session {
	session := sessionFromStore(c)
	// If Not Not New
	if !session.NotNew {
		var err error
		session, err = authTokenLogin(c)
		if err != nil {
			if err != sql.ErrNoRows && err != token.ErrAuthCookie {
				logger.Error(err)
			}
			// empty session
			session = new(Session)
		}
	}
	return session
}

// sessionFromStore get user session from session store
func sessionFromStore(c *fiber.Ctx) *Session {
	session := new(Session)
	store := g.Session.Get(c)
	sessionBytes, err := store.GetBytes(storeName)
	if err != nil {
		return session
	}

	_, err = session.UnmarshalMsg(sessionBytes)
	logger.Error(err)

	// Set NotNew because token was fetched from store
	session.NotNew = true

	if !session.LoggedIn {
		return session
	}

	if err := session.checkToken(c); err != nil {
		// default session values
		session = new(Session)
	}
	return session
}

// CanDecrypt can session decrypt
func (s *Session) CanDecrypt() bool {
	return s.PrivateKey.Armored != ""
}

var (
	errToken = errors.New("session/session: Error Token")
)

func (s *Session) checkToken(c *fiber.Ctx) error {
	if len(s.LoginToken) == 0 || s.UserID == 0 {
		return errToken
	}

	pkSalt, err := key.GetPrivateKeySalt(s.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
		}
		return errToken
	}

	if subtle.ConstantTimeCompare(s.LoginToken, genLoginToken(c, []byte(pkSalt))) == 0 {
		return errToken
	}

	return nil
}

func genLoginToken(c *fiber.Ctx, userData []byte) []byte {
	var token bytes.Buffer
	token.Write(userData)
	token.WriteString(c.Get("user-agent"))
	hash := sha512.Sum512_256(token.Bytes())
	return hash[:]
}
