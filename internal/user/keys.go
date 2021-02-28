package user

import (
	"errors"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/backupkey"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/pkg/aesgcm"
	"github.com/allocamelus/allocamelus/pkg/argon2id"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

var (
	// ErrGeneratingKeys Error for when pgp key generation or encryption fails
	ErrGeneratingKeys = errors.New("user/user: Error generating/encrypting user keys")
)

// GenerateKeys from User
//
// u.UniqueName must be != "" or pgp.NewKey will throw error
func (u *User) GenerateKeys(password string) error {
	if u.UniqueName == "" {
		logger.Error(errors.New("user/keys: Error Empty namefound"))
		return ErrGeneratingKeys
	}

	err := u.generateKeys(password)
	if logger.Error(err) {
		return ErrGeneratingKeys
	}
	return nil
}

func (u *User) generateKeys(password string) error {
	cost := g.Config.Argon2Cost

	// Make sure keyLen is 32 bytes
	cost.KeyLen = 32

	// Hash Password
	hashedObj := argon2id.Hash(password, cost)
	// Check KeyLen
	if len(hashedObj.Key) != 32 {
		return errors.New("user/user: Error password hash is != to len 32")
	}

	// Generate pgp key pair
	privateKey, err := pgp.NewKey(u.UniqueName, "@"+g.Config.Site.Domain)
	if err != nil {
		return err
	}

	u.PrivateKeySalt = hashedObj.ToStringNoKey()
	// Encrypt privateKey
	u.PrivateKey = aesgcm.EncryptBase64(hashedObj.Key, []byte(privateKey.Armored))

	// Generate backup/recovery key
	backupKey, encodedBackupKey := backupkey.Create()
	u.BackupKey = aesgcm.EncryptBase64(backupKey, []byte(privateKey.Armored))
	u.encodedBackupKey = encodedBackupKey

	return nil
}
