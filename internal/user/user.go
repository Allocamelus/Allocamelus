package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/jdinabox/goutils/logger"
)

// Perms permissions
type Perms int64

const (
	// Post perm allows user to post
	Post Perms = 1 << iota
	// Comment perm allows user to comment
	Comment
	// UploadMedia perm allows user to upload media
	UploadMedia
)

//go:generate msgp

// User Struct
type User struct {
	ID         int64  `msg:"id" json:"id"`
	UniqueName string `msg:"uniqueName" json:"uniqueName"`
	Name       string `msg:"name" json:"name"`
	Email      string `msg:"email" json:"email"`
	Avatar     bool   `msg:"avatar" json:"avatar"`
	Bio        string `msg:"bio" json:"bio"`
	Likes      int64  `msg:"likes" json:"likes"`
	Roles      Perms  `msg:"roles,omitempty" json:"-"`
	Created    int64  `msg:"created" json:"created"`
	Password   string `msg:"password,omitempty" json:"-"`
	PublicKey  string `msg:"publicKey,omitempty" json:"publicKey,omitempty"`
	// Salt used in Argon2id to derive encryption key
	PrivateSalt string `msg:"privateSalt,omitempty" json:"-"`
	PrivateKey  string `msg:"privateKey,omitempty" json:"-"`
	// Backup Private Key
	BackupKey string `msg:"backupKey,omitempty" json:"-"`
}

// New user with generated keys and hashed password
func New(plainTxtPassword string) *User {
	return &User{}
}

var preCreate *sql.Stmt

func initCreate(p data.Prepare) {
	preCreate = p(`INSERT INTO Users (uniqueName, name, email, avatar, bio, password, permissions, created, publicKey, privateSalt, privateKey, backupKey)
		VALUES (?, ?, ?, ?, '', ?, ?, ?, ?, ?, ?, ?)`)
}

// Create create new user in database
// 	returns user Id & backup Key on success
func Create(user *User) (id int64, backupKey string, err error) {

	// Insert user into database
	r, err := preCreate.Exec(
		user.UniqueName, user.Name,
		user.Email, user.Avatar,
		user.Password, user.Roles,
		user.Created, user.PublicKey,
		user.PrivateSalt, user.PrivateKey,
		user.BackupKey,
	)
	if err != nil {
		return
	}

	id, err = r.LastInsertId()
	// err not expected here with proper setup
	logger.Error(err)

	return id, "", nil
}
