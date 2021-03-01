//go:generate msgp

package token

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"k8s.io/klog/v2"
)

// Types for user
type Types int8

const (
	// Email token for email verification
	Email Types = iota
	// Reset token for password resets
	Reset
	// Auth token for user authentication
	Auth
)

const (
	selectorBytes  int = 9
	tokenBytes     int = 32
	tokenCodeBytes int = 8
	// 1 hours
	emailMaxLife time.Duration = time.Hour * 1
	// 1 hours
	resetMaxLife time.Duration = time.Hour * 1
	// 30 Days
	authMaxLife time.Duration = time.Hour * 24 * 30
)

// TypeMaxLife returns max life for type
func TypeMaxLife(t Types) (maxLife time.Duration) {
	switch t {
	case Email:
		maxLife = emailMaxLife
	case Reset:
		maxLife = resetMaxLife
	case Auth:
		maxLife = authMaxLife
	}
	return
}

// Token for user
type Token struct {
	ID         int64  `msg:"id"`
	UserID     int64  `msg:"userId"`
	Type       Types  `msg:"type"`
	Selector   string `msg:"selector"`
	Token      string `msg:"token"`
	TokenHash  string `msg:"tokenHash"`
	Expiration int64  `msg:"expiration"`
}

// New Token
func New(t Types, userID int64) *Token {
	token := new(Token)
	token.UserID = userID
	token.Type = t
	token.generatePair()
	token.Expiration = time.Now().Add(TypeMaxLife(t)).Unix()
	return token
}

// NewAndInsert Create new token and insert
func NewAndInsert(t Types, userID int64) (*Token, error) {
	token := New(t, userID)
	if err := token.Insert(); err != nil {
		return nil, err
	}
	return token, nil
}

var (
	preSelectorExist *sql.Stmt
	preInsert        *sql.Stmt
)

// Init prepared statements
func Init(p data.Prepare) {
	preSelectorExist = p(`SELECT EXISTS(SELECT * FROM UserTokens WHERE selector = ?)`)
	preInsert = p(`INSERT INTO UserTokens (userId, tokenType, selector, token, expiration) VALUES (?, ?, ?, ?, ?)`)
}

// Insert token into database
func (t *Token) Insert() error {
	_, err := preInsert.Exec(
		t.UserID, t.Type,
		t.Selector, t.TokenHash,
		t.Expiration,
	)
	return err
}

func (t *Token) generatePair() {
	t.Selector = genSelector()
	t.Token, t.TokenHash = genTokenPair(t.Type)
}

func genSelector() (selector string) {
	for {
		selector = random.StringBase64(selectorBytes)

		var exist bool
		err := preSelectorExist.QueryRow(selector).Scan(&exist)
		if err != sql.ErrNoRows {
			logger.Error(err)
		}

		if exist {
			klog.Info("Token Selector Collision found:", selector)
		} else {
			break
		}
	}
	return
}

func genTokenPair(t Types) (token, tokenHash string) {
	token = random.StringBase64(tokenBytes)
	tokenHash = hashToken(token)
	return
}

func hashToken(token string) string {
	hashedToken := sha512.Sum512_256([]byte(token))
	return base64.RawStdEncoding.EncodeToString(hashedToken[:])
}
