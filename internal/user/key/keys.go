package key

import (
	"context"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
)

// keyRecoveryTime 30 Days
const keyRecoveryTime time.Duration = time.Hour * 24 * 30

// GetPublicKeys get user's encrypted publicKey
func GetPublicKeys(userID int64) (publicKeys []*Public, err error) {
	notBefore := time.Now().Add(-keyRecoveryTime).Unix()
	rows, err := g.Data.Queries.GetUserPublicKeys(context.Background(), db.GetUserPublicKeysParams{Userid: userID, Replaced: notBefore})
	if err != nil {
		return
	}

	for _, r := range rows {
		key := new(Public)
		key.UserID = userID
		key.ID = r.Userkeyid
		key.PublicArmored = pgp.PublicKey(r.Publicarmored)

		publicKeys = append(publicKeys, key)
	}
	return
}
