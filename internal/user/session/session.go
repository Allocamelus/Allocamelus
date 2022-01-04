//go:generate msgp

package session

import (
	"bytes"
	"crypto/sha512"
	"crypto/subtle"
	"database/sql"
	"errors"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/internal/user/perms"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

// Session user session struct
type Session struct {
	LoggedIn   bool        `msg:"loggedIn" json:"loggedIn"`
	UserID     int64       `msg:"userId" json:"userId"`
	Perms      perms.Perms `msg:"perms" json:"perms"`
	LoginToken []byte      `msg:"loginToken" json:"-"`
	NotNew     bool        `msg:"notNew"  json:"notNew"`
}

const storeName = "session"

// New new session
func New(c *fiber.Ctx, userID int64) (*Session, error) {
	var err error
	session := new(Session)
	session.LoggedIn = true
	session.UserID = userID
	if err != nil {
		return nil, err
	}

	pkSalt, err := key.GetSalt(userID)
	if err != nil {
		return nil, err
	}

	perms, err := perms.Get(userID)
	if err != nil {
		return nil, err
	}

	session.Perms = perms
	session.LoginToken = genLoginToken(c, []byte(pkSalt))
	return session, nil
}

// ToContext set user session to context
func ToContext(c *fiber.Ctx) {
	c.Locals(storeName, Get(c))
}

// Context get user session from fiber context
func Context(c *fiber.Ctx) *Session {
	session := c.Locals(storeName)
	if session != nil {
		return session.(*Session)
	}
	return &Session{}
}

// LoggedIn User
func LoggedIn(c *fiber.Ctx) bool {
	return Context(c).LoggedIn
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

// Get get session from store
// attempt auth login on New session
func Get(c *fiber.Ctx) *Session {
	session := fromStore(c)
	// If Not Not New
	if !session.NotNew {
		var err error
		session, err = authTokenLogin(c)
		if err != nil {
			if err != sql.ErrNoRows &&
				err != token.ErrAuthCookie &&
				err != token.ErrInvalid {
				logger.Error(err)
			} else if klog.V(5).Enabled() {
				logger.Error(err)
			}
			// empty session
			session = new(Session)
		}
	}
	return session
}

// fromStore get user session from session store
func fromStore(c *fiber.Ctx) *Session {
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

var (
	errToken = errors.New("session/session: Error Token")
)

func (s *Session) checkToken(c *fiber.Ctx) error {
	if len(s.LoginToken) == 0 || s.UserID == 0 {
		return errToken
	}

	pkSalt, err := key.GetSalt(s.UserID)
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

func authTokenLogin(c *fiber.Ctx) (*Session, error) {
	userToken, err := token.GetAuth(c)
	if err != nil {
		return nil, err
	}

	session, err := New(c, userToken.UserID)
	if err != nil {
		return nil, err
	}
	if err := session.ToStore(c); err != nil {
		return nil, err
	}
	return session, nil
}

func genLoginToken(c *fiber.Ctx, userData []byte) []byte {
	var token bytes.Buffer
	token.Write(userData)
	token.WriteString(c.Get("user-agent"))
	hash := sha512.Sum512_256(token.Bytes())
	return hash[:]
}
