package key

import (
	"errors"

	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/pkg/aesgcm"
	"github.com/allocamelus/allocamelus/pkg/argon2id"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

// GetAndDecryptPK get and decrypt private key
//	return string (private key) error (err)
func GetAndDecryptPK(userID int64, password string) (*pgp.PrivateKey, error) {
	pk := new(pgp.PrivateKey)
	cryptPK, cost, err := GetPrivKeyAndSalt(userID)
	if err != nil {
		return nil, err
	}

	costObj, _, err := argon2id.Parse(cost)
	// Should not happen log just incase
	if logger.Error(err) {
		return nil, errors.New("user/keys: Error Parsing password cost")
	}

	passwordObj := argon2id.HashSalt(password, costObj.Salt, costObj.Cost)

	pkArmored, err := aesgcm.DecryptBase64(passwordObj.Key, cryptPK)
	if err != nil {
		return nil, ErrDecryptingKey
	}

	pk.Armored = string(pkArmored)
	return pk, nil
}
