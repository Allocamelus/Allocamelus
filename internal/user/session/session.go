//go:generate msgp

package session

import (
	"bytes"
	"crypto/sha512"
	"crypto/subtle"
	"database/sql"
	"errors"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// Session user session struct
type Session struct {
	LoggedIn   bool       `msg:"loggedIn"`
	UserID     int64      `msg:"userId"`
	Perms      user.Perms `msg:"perms"`
	LoginToken []byte     `msg:"loginToken"`
}

const storeName = "session"

var errInvalidID = errors.New("session/session: Error Invalid UserID")

// NewFromID new session from user id
func NewFromID(c *fiber.Ctx, userID int64) (*Session, error) {
	session := new(Session)
	session.LoggedIn = true
	session.UserID = userID

	pkSalt, err := user.GetPrivateKeySalt(userID)
	if logger.Error(err) {
		return nil, errInvalidID
	}
	perms, err := user.GetPerms(userID)
	if logger.Error(err) {
		return nil, errInvalidID
	}

	session.Perms = perms
	session.LoginToken = genLoginToken(c, []byte(pkSalt))
	return session, nil
}

// ToStore set user session to session store
func (s *Session) ToStore(c *fiber.Ctx) {
	if s != nil {
		store := g.Session.Get(c)
		store.Regenerate()
		store.Set(storeName, *s)
	} else {
		logger.Error(errors.New("session/session: nil *Session"))
	}
}

// FromStore get user session from session store
func FromStore(c *fiber.Ctx) *Session {
	session := Session{}
	store := g.Session.Get(c)
	sessionInter, err := store.Get(storeName)
	if err != nil {
		return &session
	}

	session = sessionInter.(Session)
	if !session.LoggedIn {
		return &session
	}

	if err := session.checkToken(c); err != nil {
		// default session values
		session = Session{}
	}
	return &session
}

var (
	errToken = errors.New("session/session: Error Token")
)

func (s *Session) checkToken(c *fiber.Ctx) error {
	if len(s.LoginToken) == 0 || s.UserID == 0 {
		return errToken
	}

	pkSalt, err := user.GetPrivateKeySalt(s.UserID)
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
