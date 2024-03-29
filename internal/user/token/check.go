package token

import (
	"crypto/subtle"
	"database/sql"
	_ "embed"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/pkg/logger"
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
	//go:embed sql/get.sql
	qGet   string
	preGet *sql.Stmt
	//go:embed sql/delete.sql
	qDelete   string
	preDelete *sql.Stmt
	//go:embed sql/delByUIDAndType.sql
	qDelByUIDAndType   string
	preDelByUIDAndType *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGet, qGet)
	data.PrepareQueuer.Add(&preDelete, qDelete)
	data.PrepareQueuer.Add(&preDelByUIDAndType, qDelByUIDAndType)
}

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
	err := preGet.QueryRow(selector).Scan(&token.ID, &token.UserID, &token.Type, &token.TokenHash, &token.Expiration)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
		}
		return nil, ErrInvalidSelector
	}

	return token, nil
}

// Delete token from database
//
// Token SHOULD NOT be used after successful delete
func (t *Token) Delete() error {
	_, err := preDelete.Exec(t.ID)
	if err != nil {
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

	if subtle.ConstantTimeEq(int32(t.Type), int32(ty)) == 0 ||
		subtle.ConstantTimeEq(int32(t.UserID), int32(userID)) == 0 ||
		subtle.ConstantTimeEq(int32(t.UserID>>32), int32(userID>>32)) == 0 ||
		subtle.ConstantTimeCompare([]byte(t.TokenHash), []byte(hashToken(token))) == 0 {
		return ErrInvalidToken
	}

	return nil
}

// IsExpired has token Expired
func (t *Token) IsExpired() bool {
	return t.Expiration < time.Now().Unix()
}
