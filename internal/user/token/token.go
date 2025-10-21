//go:generate msgp

package token

import (
	"context"
	"crypto/sha512"
	_ "embed"
	"encoding/base64"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"github.com/jackc/pgx/v5"
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
	selectorBytes int = 9
	tokenBytes    int = 32
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
	token      string
	TokenHash  string `msg:"tokenHash"`
	Created    int64  `msg:"created"`
	Expiration int64  `msg:"expiration"`
}

// New Token
func New(t Types, userID int64) *Token {
	token := new(Token)
	token.UserID = userID
	token.Type = t
	token.generatePair()
	token.Created = time.Now().Unix()
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

// Insert token into database
func (t *Token) Insert() error {
	return g.Data.Queries.InsertUserToken(context.Background(), db.InsertUserTokenParams{
		Userid: t.UserID, Tokentype: int16(t.Type), Selector: t.Selector,
		Token: t.TokenHash, Created: t.Created, Expiration: t.Expiration,
	})
}

// GetToken return token string
func (t *Token) GetToken() string {
	return t.token
}

func (t *Token) generatePair() {
	t.Selector = genSelector()
	t.token, t.TokenHash = genTokenPair()
}

func genSelector() (selector string) {
	for {
		selector = random.StringBase64(selectorBytes)

		exist, err := g.Data.Queries.UserTokenSelectorExist(context.Background(), selector)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
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

func genTokenPair() (token, tokenHash string) {
	token = random.StringBase64(tokenBytes)
	tokenHash = hashToken(token)
	return
}

func hashToken(token string) string {
	hashedToken := sha512.Sum512_256([]byte(token))
	return base64.RawStdEncoding.EncodeToString(hashedToken[:])
}
