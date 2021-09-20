//go:generate msgp

package key

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
)

// Key user key struct
type Key struct {
	ID      int64 `msg:"id"`
	userID  int64 `msg:"userId"`
	created int64 `msg:"created"`
	//replaced  int64         `msg:"replaced"`
	PublicKey pgp.PublicKey `msg:"publicKey,omitempty"`
	// Salt used in Argon2id to derive encryption key
	PrivateKeySalt string `msg:"privateKeySalt,omitempty"`
	PrivateKey     string `msg:"encPrivateKey,omitempty"`
	//armoredPrivateKey pgp.PrivateKey `msg:"privateKey,omitempty"`
	// Backup PrivateKey encrypted with encodedBackupKey
	BackupKey string `msg:"backupKey,omitempty"`
	// Hash of encodedBackupKey
	RecoveryKeyHash string `msg:"recoveryKeyHash,omitempty"`
	// Encoded Key for encrypting BackupKey
	encodedBackupKey string
}

var (
	// ErrGeneratingKeys Error for when pgp key generation or encryption fails
	ErrGeneratingKeys = errors.New("user/user: Error generating/encrypting user keys")
	// ErrDecryptingKey Error Decrypting Key
	ErrDecryptingKey     = errors.New("user/user: Error Decrypting Key")
	preKeyID             *sql.Stmt
	prePublicKey         *sql.Stmt
	preActivePublicKeys  *sql.Stmt
	prePrivateKey        *sql.Stmt
	prePrivateKeySalt    *sql.Stmt
	prePrivKeyAndSalt    *sql.Stmt
	preBackupKey         *sql.Stmt
	preUpdateKeyReplaced *sql.Stmt
	preInsertKey         *sql.Stmt
)

// keyRecoveryTime 30 Days
const keyRecoveryTime time.Duration = time.Hour * 24 * 30

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
	prePublicKey = keySelectQueryBuilder(p, "userKeyId", "publicKey")
	preActivePublicKeys = p(`SELECT userKeyId, publicKey FROM UserKeys WHERE userId = ? AND (replaced > ? OR replaced = 0) ORDER BY userKeyId DESC`)
	prePrivateKey = keySelectQueryBuilder(p, "privateKey")
	prePrivateKeySalt = keySelectQueryBuilder(p, "keySalt")
	prePrivKeyAndSalt = keySelectQueryBuilder(p, "privateKey", "keySalt")
	preBackupKey = keySelectQueryBuilder(p, "backupKey")
	preUpdateKeyReplaced = p(`UPDATE UserKeys SET replaced = ? WHERE userKeyId = ?`)
	preInsertKey = p(`INSERT INTO UserKeys (userId, created, publicKey, keySalt, privateKey, recoveryKeyHash, backupKey)
		VALUES (?, ?, ?, ?, ?, ?, ?)`)
}

// NewPair Generate Key pair from userId and password
func NewPair(userID int64, password string) (*Key, error) {
	key := new(Key)
	key.userID = userID
	// generate new keys
	err := key.Generate(password)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// InsertNew Create and insert User Key pair
func InsertNew(userID int64, password string) (*Key, error) {
	kp, err := NewPair(userID, password)
	if err != nil {
		return nil, err
	}
	err = kp.Insert()
	if err != nil {
		return nil, err
	}
	return kp, nil
}

// UpdateKey for user
func (k *Key) UpdateKey() error {
	keyID, err := GetKeyID(k.userID)
	if err != nil {
		return err
	}
	// Update old key's replaced time
	_, err = preUpdateKeyReplaced.Exec(time.Now().Unix(), keyID)
	if err != nil {
		return err
	}

	return k.Insert()
}

// Insert key into database
func (k *Key) Insert() error {
	_, err := preInsertKey.Exec(
		k.userID, k.created,
		k.PublicKey.Armored, k.PrivateKeySalt,
		k.PrivateKey, k.RecoveryKeyHash,
		k.BackupKey,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetKeyID get user's Key id
func GetKeyID(userID int64) (keyID int64, err error) {
	err = preKeyID.QueryRow(userID).Scan(&keyID)
	return
}

// GetPublicKey get user's encrypted publicKey
func GetPublicKey(userID int64) (pk *Key, err error) {
	pk.userID = userID
	err = prePublicKey.QueryRow(userID).Scan(&pk.ID, &pk.PublicKey.Armored)
	return
}

// GetPublicKeys get user's encrypted publicKey
func GetPublicKeys(userID int64) (publicKeys []*Key, err error) {
	notBefore := time.Now().Add(-keyRecoveryTime).Unix()
	rows, err := preActivePublicKeys.Query(userID, notBefore)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		key := new(Key)
		key.userID = userID
		err = rows.Scan(&key.ID, &key.PublicKey.Armored)
		if err != nil {
			return
		}
		publicKeys = append(publicKeys, key)
	}
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
func (k *Key) GetEncodedBackupKey() string {
	return k.encodedBackupKey
}
