package user

import (
	"database/sql"
	"errors"

	"github.com/allocamelus/allocamelus/internal/data"
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
	// ErrDecryptingKey Error Decrypting Key
	ErrDecryptingKey  = errors.New("user/user: Error Decrypting Key")
	prePublicKey      *sql.Stmt
	prePrivateKey     *sql.Stmt
	prePrivateKeySalt *sql.Stmt
)

func initKeys(p data.Prepare) {
	prePublicKey = p(`SELECT publicKey FROM Users WHERE userId = ? LIMIT 1`)
	prePrivateKey = p(`SELECT privateKey FROM Users WHERE userId = ? LIMIT 1`)
	prePrivateKeySalt = p(`SELECT privateKeySalt FROM Users WHERE userId = ? LIMIT 1`)
}

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

// GetAndDecryptPK get and decrypt private key
//	return string (private key) error (err)
func GetAndDecryptPK(userID int64, password string) (*pgp.PrivateKey, error) {
	pk := new(pgp.PrivateKey)
	cryptPK, err := GetPrivateKey(userID)
	if err != nil {
		return nil, err
	}
	cost, err := GetPrivateKeySalt(userID)
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

// GetPublicKey get user's encrypted publicKey
func GetPublicKey(userID int64) (pk pgp.PublicKey, err error) {
	err = prePublicKey.QueryRow(userID).Scan(&pk.Armored)
	return
}

// GetPrivateKey get user's encrypted privateKey
func GetPrivateKey(userID int64) (pk string, err error) {
	err = prePrivateKey.QueryRow(userID).Scan(&pk)
	return
}

// GetPrivateKeySalt get user's privateKeySalt
func GetPrivateKeySalt(userID int64) (pkSalt string, err error) {
	err = prePrivateKeySalt.QueryRow(userID).Scan(&pkSalt)
	return
}

// GetEncodedBackupKey returns user.encodedBackupKey
func (u *User) GetEncodedBackupKey() string {
	return u.encodedBackupKey
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

	publicKey, err := privateKey.PublicKey()
	if err != nil {
		return err
	}

	u.PublicKey = publicKey.Armored

	u.PrivateKeySalt = hashedObj.ToStringNoKey()
	// Encrypt privateKey
	u.PrivateKey = aesgcm.EncryptBase64(hashedObj.Key, []byte(privateKey.Armored))

	// Generate backup/recovery key
	backupKey, encodedBackupKey := backupkey.Create()
	u.BackupKey = aesgcm.EncryptBase64(backupKey, []byte(privateKey.Armored))
	u.encodedBackupKey = encodedBackupKey

	return nil
}
