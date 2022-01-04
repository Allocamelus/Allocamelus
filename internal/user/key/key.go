//go:generate msgp

package key

import (
	"database/sql"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
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

func init() {
	data.PrepareQueuer.Add(&preInsertKey, qInsertKey)
}

var (
	//go:embed sql/insertKey.sql
	qInsertKey   string
	preInsertKey *sql.Stmt
)

func (k *Private) Insert() error {
	_, err := preInsertKey.Exec(
		k.UserID,
		k.Created,
		k.AuthKeyHash,
		k.AuthKeySalt,
		k.PublicArmored,
		k.PrivateArmored,
		k.RecoveryKeyHash,
		k.RecoveryArmored,
	)
	return err
}

func init() {
	data.PrepareQueuer.Add(&preGetPrivateArmored, qGetPrivateArmored)
}

var (
	//go:embed sql/getPrivateArmored.sql
	qGetPrivateArmored   string
	preGetPrivateArmored *sql.Stmt
)

func GetPrivateArmored(userID int64) (privateArmored pgp.PrivateKey, err error) {
	err = preGetPrivateArmored.QueryRow(userID).Scan(&privateArmored)
	return
}
