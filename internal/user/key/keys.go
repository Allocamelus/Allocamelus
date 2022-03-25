package key

import (
	"database/sql"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
)

func init() {
	data.PrepareQueuer.Add(&preGetPublicKeys, qGetPublicKeys)
}

var (
	//go:embed sql/getPublicKeys.sql
	qGetPublicKeys   string
	preGetPublicKeys *sql.Stmt
)

// keyRecoveryTime 30 Days
const keyRecoveryTime time.Duration = time.Hour * 24 * 30

// GetPublicKeys get user's encrypted publicKey
func GetPublicKeys(userID int64) (publicKeys []*Public, err error) {
	notBefore := time.Now().Add(-keyRecoveryTime).Unix()
	rows, err := preGetPublicKeys.Query(userID, notBefore)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		key := new(Public)
		key.UserID = userID
		err = rows.Scan(&key.ID, &key.PublicArmored)
		if err != nil {
			return
		}
		publicKeys = append(publicKeys, key)
	}
	return
}
