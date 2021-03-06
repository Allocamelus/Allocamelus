//go:generate msgp

package user

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/user/avatar"
	"github.com/allocamelus/allocamelus/internal/user/key"
)

// Perms permissions
type Perms int64

// User Types
type Types int8

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

type FollowStruct struct {
	Following bool `msg:"-" json:"following"`
	Requested bool `msg:"-" json:"requested"`
}

// User Struct
type User struct {
	ID          int64        `msg:"id" json:"id"`
	UserName    string       `msg:"userName" json:"userName"`
	Name        string       `msg:"name" json:"name"`
	Email       string       `msg:"email" json:"email,omitempty"`
	Avatar      bool         `msg:"avatar" json:"avatar"`
	AvatarUrl   string       `msg:"-" json:"avatarUrl,omitempty"`
	Bio         string       `msg:"bio" json:"bio,omitempty"`
	SelfFollow  FollowStruct `msg:"-" json:"selfFollow,omitempty"`
	UserFollow  FollowStruct `msg:"-" json:"userFollow,omitempty"`
	Followers   int64        `msg:"followers" json:"followers"`
	Type        Types        `msg:"type" json:"type"`
	Permissions Perms        `msg:"permissions" json:"-"`
	Created     int64        `msg:"created" json:"created,omitempty"`
}

// New user
func New(userName, name, email string) *User {
	user := new(User)
	user.UserName = userName
	user.Name = name
	user.Email = email
	user.Permissions = DefaultPerms
	user.Created = time.Now().Unix()
	return user
}

var (
	preInsert    *sql.Stmt
	preGetPublic *sql.Stmt
)

func initUser(p data.Prepare) {
	preInsert = p(`INSERT INTO Users (userName, name, email, bio, type, permissions, created)
		VALUES (?, '', ?, '', ?, ?, ?)`)
	preGetPublic = p(`SELECT userName, name, bio, type, created FROM Users WHERE userId = ? LIMIT 1`)
}

// Insert new user into database
// 	returns nil and sets user.ID on success
func (u *User) Insert() error {
	// Insert user into database
	r, err := preInsert.Exec(
		u.UserName, u.Email, u.Type,
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

// GetPublic user info for session user
// TODO: Cache
func GetPublic(s *Session, userID int64) (user User, err error) {
	user.ID = userID
	err = preGetPublic.QueryRow(userID).Scan(&user.UserName, &user.Name, &user.Bio, &user.Type, &user.Created)
	if err != nil {
		return
	}

	user.AvatarUrl, err = avatar.GetUrl(userID)
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
		err = nil
	} else {
		user.Avatar = true
	}

	if s.LoggedIn {
		user.SelfFollow, err = Following(s.UserID, userID)
		if err != nil {
			return
		}
		user.UserFollow, err = Following(userID, s.UserID)
		if err != nil {
			return
		}
	}

	user.Followers, err = Followers(userID)
	return
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
