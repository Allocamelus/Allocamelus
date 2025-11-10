package token

import (
	"context"
	"crypto/subtle"
	_ "embed"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jackc/pgx/v5"
)

var (
	// ErrInvalidSelector Error Invalid Selector
	ErrInvalidSelector = errors.New("token/check: Error Invalid Selector")
	// ErrInvalidToken Error Invalid Token
	ErrInvalidToken = errors.New("token/check: Error Invalid Token")
	// ErrExpiredToken Error Expired Token
	ErrExpiredToken = errors.New("token/check: Error Expired Token")
	// ErrInvalid Error Invalid
	ErrInvalid = errors.New("token/check: Error Invalid")
)

// Check Selector Token and Type
func Check(selector, token string, t Types) (*Token, error) {
	tkn, err := Get(selector)
	if err != nil {
		return nil, ErrInvalid
	}

	if err := tkn.Check(token, tkn.UserID, t); err != nil {
		if err == ErrExpiredToken {
			return nil, err
		}
		return nil, ErrInvalid
	}
	return tkn, nil
}

// CheckWithID Selector Token UserId and Type
func CheckWithID(selector, token string, userID int64, t Types) (*Token, error) {
	tkn, err := Get(selector)
	if err != nil {
		return nil, ErrInvalid
	}

	if err := tkn.Check(token, userID, t); err != nil {
		if err == ErrExpiredToken {
			return nil, err
		}
		return nil, ErrInvalid
	}
	return tkn, nil
}

// Get Token
func Get(selector string) (*Token, error) {
	if selector == "" {
		return nil, ErrInvalidSelector
	}
	token := new(Token)
	t, err := g.Data.Queries.GetUserToken(context.Background(), selector)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			logger.Error(err)
		}
		return nil, ErrInvalidSelector
	}

	token.ID = t.Usertokenid
	token.UserID = t.Userid
	token.Type = Types(t.Tokentype)
	token.TokenHash = t.Token
	token.Expiration = t.Expiration

	return token, nil
}

// Delete token from database
//
// Token SHOULD NOT be used after successful delete
func (t *Token) Delete() error {
	if err := g.Data.Queries.DeleteUserToken(context.Background(), t.ID); err != nil {
		return err
	}
	// unset Token
	t = nil
	return nil
}

// Check check if token is valid
func (t *Token) Check(token string, userID int64, ty Types) error {
	if t.IsExpired() {
		logger.Error(t.Delete())
		return ErrExpiredToken
	}

	if t.TokenHash == "" || token == "" {
		return ErrInvalidToken
	}

	if compare.EqualInt(int8(t.Type), int8(ty)) ||
		compare.EqualInt(t.UserID, userID) ||
		subtle.ConstantTimeCompare([]byte(t.TokenHash), []byte(hashToken(token))) == 0 {
		return ErrInvalidToken
	}

	return nil
}

// IsExpired has token Expired
func (t *Token) IsExpired() bool {
	return t.Expiration < time.Now().Unix()
}
