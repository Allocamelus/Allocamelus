package token

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/gofiber/fiber/v2"
)

const (
	authNamePostfix = "auth"
	authValueSep    = "$"
)

// ErrAuthCookie Invalid Auth Cookie
var ErrAuthCookie = errors.New("token/auth: Error Invalid Auth Cookie Value")

// SetAuth create, insert and add to cookies auth token
func SetAuth(c *fiber.Ctx, userID int64) error {
	token, err := NewAndInsert(Auth, userID)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:     g.Config.Cookie.PreFix + authNamePostfix,
		Value:    token.Selector + authValueSep + token.token,
		Domain:   "." + g.Config.Site.Domain,
		Expires:  time.Now().Add(authMaxLife),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
	})
	return nil
}

// GetAuth get and check auth token
//
//	return *Token || error
func GetAuth(c *fiber.Ctx) (*Token, error) {
	cookie := c.Cookies(g.Config.Cookie.PreFix+authNamePostfix, "")
	// check if cookie len(value) is at least
	// selectorMinLen + authValueSepLen + tokenMinLen
	if len(cookie) < selectorBytes+len(authValueSep)+tokenBytes {
		unsetAuthCookie(c)
		return nil, ErrAuthCookie
	}
	cookieSlice := strings.Split(cookie, authValueSep)
	if len(cookieSlice) != 2 {
		unsetAuthCookie(c)
		return nil, ErrAuthCookie
	}

	token, err := Check(cookieSlice[0], cookieSlice[1], Auth)
	if err != nil {
		unsetAuthCookie(c)
		return nil, err
	}

	return token, nil
}

// DeleteAuth token if exist
func DeleteAuth(c *fiber.Ctx) error {
	// Check and get token
	token, err := GetAuth(c)
	if err != nil {
		return err
	}
	unsetAuthCookie(c)
	return token.Delete()
}

// DeleteAuthByID user's Auth tokens from database
func DeleteAuthByID(userID int64) error {
	return g.Data.Queries.DeleteUserTokenByUIDAndType(context.Background(), db.DeleteUserTokenByUIDAndTypeParams{Userid: userID, Tokentype: int16(Auth)})
}

func unsetAuthCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     g.Config.Cookie.PreFix + authNamePostfix,
		Value:    "",
		Domain:   "." + g.Config.Site.Domain,
		Expires:  time.Now(),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
	})
}
