//go:generate msgp

package user

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/user/key"
)

// Perms permissions
type Perms int64

// Session user session struct
type Session struct {
	LoggedIn   bool           `msg:"loggedIn" json:"loggedIn"`
	UserID     int64          `msg:"userId" json:"userId"`
	UserName   string         `msg:"userName" json:"userName"`
	Perms      Perms          `msg:"perms" json:"perms"`
	PrivateKey pgp.PrivateKey `msg:"privateKey" json:"-"`
	LoginToken []byte         `msg:"loginToken" json:"-"`
	NotNew     bool           `msg:"notNew"  json:"notNew"`
}

// User Struct
type User struct {
	ID          int64  `msg:"id" json:"id"`
	UserName    string `msg:"userName" json:"userName"`
	Name        string `msg:"name" json:"name"`
	Email       string `msg:"email" json:"email,omitempty"`
	Avatar      bool   `msg:"avatar" json:"avatar"`
	Bio         string `msg:"bio" json:"bio,omitempty"`
	Likes       int64  `msg:"likes" json:"likes"`
	Permissions Perms  `msg:"permissions" json:"-"`
	Created     int64  `msg:"created" json:"created,omitempty"`
}

// New user
func New(userName, name, email string) *User {
	user := new(User)
	user.UserName = userName
	user.Name = name
	user.Email = email
	user.Created = time.Now().Unix()
	return user
}

var (
	preInsert    *sql.Stmt
	preGetPublic *sql.Stmt
)

func initUser(p data.Prepare) {
	preInsert = p(`INSERT INTO Users (userName, name, email, avatar, bio, permissions, created)
		VALUES (?, '', ?, 0, '', ?, ?)`)
	preGetPublic = p(`SELECT userName, name, avatar, bio, created FROM Users WHERE userId = ? LIMIT 1`)
}

// Insert new user into database
// 	returns nil and sets user.ID on success
func (u *User) Insert() error {
	// Insert user into database
	r, err := preInsert.Exec(
		u.UserName, u.Email,
		u.Permissions, u.Created,
	)
	if err != nil {
		return err
	}

	u.ID, err = r.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

// GetPublic user info
// TODO: Likes
// TODO: Cache
func GetPublic(userID int64) (User, error) {
	var u User
	u.ID = userID
	err := preGetPublic.QueryRow(userID).Scan(&u.UserName, &u.Name, &u.Avatar, &u.Bio, &u.Created)
	return u, err
}

// UpdatePassword for User
//
// Password Reset Event should be called after this
func UpdatePassword(userID int64, password string) (backupKey string, err error) {
	k, err := key.NewPair(userID, password)
	if err != nil {
		return
	}
	// update key
	err = k.UpdateKey()
	if err != nil {
		return
	}
	return k.GetEncodedBackupKey(), nil
}
