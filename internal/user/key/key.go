//go:generate msgp

package key

import (
	"context"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
)

type Public struct {
	ID            int64         `json:"id" msg:"id"`
	UserID        int64         `json:"userId" msg:"userId"`
	Created       int64         `json:"created" msg:"created"`
	Replaced      int64         `json:"replaced,omitempty" msg:"replaced"`
	PublicArmored pgp.PublicKey `json:"publicArmored" msg:"publicArmored"`
}

type Private struct {
	Public
	AuthKeyHash     string
	AuthKeySalt     string
	PrivateArmored  pgp.PrivateKey
	RecoveryKeyHash string
	RecoveryArmored pgp.PrivateKey
}

func NewPrivate() *Private {
	k := new(Private)
	k.Created = time.Now().Unix()
	return k
}

func (k *Private) Insert() error {
	return g.Data.Queries.InsertUserKey(context.Background(), db.InsertUserKeyParams{
		Userid:          k.UserID,
		Created:         k.Created,
		Authkeyhash:     k.AuthKeyHash,
		Authkeysalt:     k.AuthKeySalt,
		Publicarmored:   k.PublicArmored.ToString(),
		Privatearmored:  k.PrivateArmored.ToString(),
		Recoverykeyhash: k.RecoveryKeyHash,
		Recoveryarmored: k.RecoveryArmored.ToString(),
	})
}

func GetPrivateArmored(userID int64) (privateArmored pgp.PrivateKey, err error) {
	p, err := g.Data.Queries.GetUserPrivateArmoredKey(context.Background(), userID)
	privateArmored = pgp.PrivateKey(p)
	return
}
