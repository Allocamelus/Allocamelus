package user

import (
	"database/sql"
	"errors"
	"strings"
	"time"

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
	ErrDecryptingKey     = errors.New("user/user: Error Decrypting Key")
	preKeyID             *sql.Stmt
	prePublicKey         *sql.Stmt
	prePrivateKey        *sql.Stmt
	prePrivateKeySalt    *sql.Stmt
	prePrivKeyAndSalt    *sql.Stmt
	preBackupKey         *sql.Stmt
	preUpdateKeyReplaced *sql.Stmt
	preInsertKey         *sql.Stmt
)

func keySelectQueryBuilder(p data.Prepare, items ...string) *sql.Stmt {
	var itemsPart strings.Builder
	for i, v := range items {
		if i > 0 {
			itemsPart.WriteRune(',')
		}
		itemsPart.WriteString(" " + v)

	}
	return p(`SELECT` + itemsPart.String() + ` FROM UserKeys WHERE userId = ? ORDER BY userKeyId DESC LIMIT 1`)
}

func initKeys(p data.Prepare) {
	preKeyID = keySelectQueryBuilder(p, "userKeyId")
	prePublicKey = keySelectQueryBuilder(p, "publicKey")
	prePrivateKey = keySelectQueryBuilder(p, "privateKey")
	prePrivateKeySalt = keySelectQueryBuilder(p, "keySalt")
	prePrivKeyAndSalt = keySelectQueryBuilder(p, "privateKey", "keySalt")
	preBackupKey = keySelectQueryBuilder(p, "backupKey")
	preUpdateKeyReplaced = p(`UPDATE UserKeys SET replaced = ? WHERE userKeyId = ?`)
	preInsertKey = p(`INSERT INTO UserKeys (userId, created, publicKey, keySalt, privateKey, backupKey)
		VALUES (?, ?, ?, ?, ?, ?)`)
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

// UpdatePassword for User
//
// Password Reset Event should be called after this
func UpdatePassword(userID int64, password string) (backupKey string, err error) {
	var user User
	user.ID = userID
	// Used in key generation
	user.UniqueName, err = GetUniqueNameByID(userID)
	if err != nil {
		return
	}
	// generate new keys
	err = user.GenerateKeys(password)
	if err != nil {
		return
	}
	// update key
	err = user.UpdateKey()
	if err != nil {
		return
	}
	return user.encodedBackupKey, nil
}

// UpdateKey for user
func (u *User) UpdateKey() error {
	keyID, err := GetKeyID(u.ID)
	if err != nil {
		return err
	}
	// Update old key's replaced time
	_, err = preUpdateKeyReplaced.Exec(time.Now().Unix(), keyID)
	if err != nil {
		return err
	}

	return u.insertKey()
}

func (u *User) insertKey() error {
	_, err := preInsertKey.Exec(
		u.ID, time.Now().Unix(),
		u.PublicKey, u.PrivateKeySalt,
		u.PrivateKey, u.BackupKey,
	)
	if err != nil {
		return err
	}
	return nil
}

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

// GetKeyID get user's Key id
func GetKeyID(userID int64) (keyID, err error) {
	err = preKeyID.QueryRow(userID).Scan(&keyID)
	return
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

// GetPrivKeyAndSalt get user's privateKey and Salt
func GetPrivKeyAndSalt(userID int64) (pk string, pkSalt string, err error) {
	err = prePrivKeyAndSalt.QueryRow(userID).Scan(&pk, &pkSalt)
	return
}

// GetBackupKey get user's encrypted BackupKey
func GetBackupKey(userID int64) (bk string, err error) {
	err = preBackupKey.QueryRow(userID).Scan(&bk)
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
