package key

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/backupkey"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/pkg/aesgcm"
	"github.com/allocamelus/allocamelus/pkg/argon2id"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"golang.org/x/crypto/sha3"
)

// Generate Keys from User
func (k *Key) Generate(password string) error {
	err := k.generateKeys(password)
	if logger.Error(err) {
		return ErrGeneratingKeys
	}
	return nil
}

func (k *Key) generateKeys(password string) error {
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
	privateKey, err := pgp.NewKey(strconv.Itoa(int(k.userID)), "@"+g.Config.Site.Domain)
	if err != nil {
		return err
	}

	publicKey, err := privateKey.PublicKey()
	if err != nil {
		return err
	}

	k.PublicKey = *publicKey

	k.PrivateKeySalt = hashedObj.ToStringNoKey()
	// Encrypt privateKey
	k.PrivateKey = aesgcm.EncryptBase64(hashedObj.Key, []byte(privateKey.Armored))

	// Generate backup/recovery key
	backupKey, encodedBackupKey := backupkey.Create()
	k.BackupKey = aesgcm.EncryptBase64(backupKey, []byte(privateKey.Armored))
	k.RecoveryKeyHash = hashRecoveryKey(backupKey)
	k.encodedBackupKey = encodedBackupKey

	return nil
}

func hashRecoveryKey(rk []byte) string {
	s2 := sha512.Sum512(rk)
	s3 := sha3.Sum512(s2[:])
	return base64.RawStdEncoding.EncodeToString(s3[:])
}
