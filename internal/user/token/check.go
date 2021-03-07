package token

import (
	"crypto/subtle"
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/pkg/byteutil"
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
	ErrInvalid         = errors.New("token/check: Error Invalid")
	preGet             *sql.Stmt
	preDelete          *sql.Stmt
	preDelByUIDAndType *sql.Stmt
)

func initCheck(p data.Prepare) {
	preGet = p(`SELECT userTokenId, userId,	tokenType, token, expiration FROM UserTokens WHERE selector = ? LIMIT 1`)
	preDelete = p(`DELETE FROM UserTokens WHERE userTokenId=?`)
	preDelByUIDAndType = p(`DELETE FROM UserTokens WHERE userId=? AND tokenType = ?`)
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
	if len(selector) == 0 {
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

	if len(t.TokenHash) == 0 || len(token) == 0 {
		return ErrInvalidToken
	}

	if subtle.ConstantTimeEq(int32(t.Type), int32(ty)) == 1 {
		if subtle.ConstantTimeCompare(byteutil.Itob(int(t.UserID)), byteutil.Itob(int(userID))) == 1 {
			if subtle.ConstantTimeCompare([]byte(t.TokenHash), []byte(hashToken(token))) == 1 {
				return nil
			}
		}
	}

	return ErrInvalidToken
}

// IsExpired has token Expired
func (t *Token) IsExpired() bool {
	if t.Expiration < time.Now().Unix() {
		return true
	}
	return false
}
