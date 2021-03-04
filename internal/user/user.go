//go:generate msgp

package user

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

// Perms permissions
type Perms int64

// Session user session struct
type Session struct {
	LoggedIn   bool           `msg:"loggedIn"`
	UserID     int64          `msg:"userId"`
	Perms      Perms          `msg:"perms"`
	PrivateKey pgp.PrivateKey `msg:"privateKey"`
	LoginToken []byte         `msg:"loginToken"`
	NotNew     bool           `msg:"notNew"`
}

// User Struct
type User struct {
	ID          int64  `msg:"id" json:"id"`
	UniqueName  string `msg:"uniqueName" json:"uniqueName"`
	Name        string `msg:"name" json:"name"`
	Email       string `msg:"email" json:"email"`
	Avatar      bool   `msg:"avatar" json:"avatar"`
	Bio         string `msg:"bio" json:"bio"`
	Likes       int64  `msg:"likes" json:"likes"`
	Permissions Perms  `msg:"permissions" json:"-"`
	Created     int64  `msg:"created" json:"created"`
	PublicKey   string `msg:"publicKey,omitempty" json:"publicKey,omitempty"`
	// Salt used in Argon2id to derive encryption key
	PrivateKeySalt string `msg:"privateKeySalt,omitempty" json:"-"`
	PrivateKey     string `msg:"privateKey,omitempty" json:"-"`
	// Backup PrivateKey encrypted with encodedBackupKey
	BackupKey string `msg:"backupKey,omitempty" json:"-"`
	// Encoded Key for encrypting BackupKey
	encodedBackupKey string
}

// New user
func New(uniqueName, name, email string) *User {
	user := new(User)
	user.UniqueName = uniqueName
	user.Name = name
	user.Email = email
	user.Created = time.Now().Unix()
	return user
}

var preInsert *sql.Stmt

func initCreate(p data.Prepare) {
	preInsert = p(`INSERT INTO Users (uniqueName, name, email, avatar, bio, permissions, created, publicKey, privateKeySalt, privateKey, backupKey)
		VALUES (?, ?, ?, '0', '', ?, ?, ?, ?, ?, ?)`)
}

// Insert new user into database
// 	returns userId int64 & encodedBackupKey string on success
func (u *User) Insert() (string, error) {
	// Insert user into database
	r, err := preInsert.Exec(
		u.UniqueName, u.Name,
		u.Email, u.Permissions,
		u.Created, u.PublicKey,
		u.PrivateKeySalt, u.PrivateKey,
		u.BackupKey,
	)
	if err != nil {
		return "", err
	}

	u.ID, err = r.LastInsertId()
	// err not expected here with proper setup
	logger.Error(err)

	return u.GetEncodedBackupKey(), nil
}
