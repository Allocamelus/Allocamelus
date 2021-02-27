//go:generate msgp

package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/jdinabox/goutils/logger"
)

// Perms permissions
type Perms int64

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
	// Backup Private Key
	BackupKey string `msg:"backupKey,omitempty" json:"-"`
	// Encoded Key for encrypting BackupKey
	encodedBackupKey string
}

// New user with generated keys and hashed password
func New(uniqueName, name, email string) (*User, error) {
	user := new(User)
	user.UniqueName = uniqueName
	user.Name = name
	user.Email = email
	user.SetDefaultPerms()
	return user, nil
}

var preCreate *sql.Stmt

func initCreate(p data.Prepare) {
	preCreate = p(`INSERT INTO Users (uniqueName, name, email, avatar, bio, permissions, created, publicKey, privateKeySalt, privateKey, backupKey)
		VALUES (?, ?, ?, '0', '', ?, ?, ?, ?, ?, ?)`)
}

// Insert new user into database
// 	returns user Id & backup Key on success
func (u *User) Insert() (id int64, backupKey string, err error) {
	// Insert user into database
	r, err := preCreate.Exec(
		u.UniqueName, u.Name,
		u.Email, u.Permissions,
		u.Created, u.PublicKey,
		u.PrivateKeySalt, u.PrivateKey,
		u.BackupKey,
	)
	if err != nil {
		return
	}

	id, err = r.LastInsertId()
	// err not expected here with proper setup
	logger.Error(err)

	return id, "", nil
}