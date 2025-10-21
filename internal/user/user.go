//go:generate msgp

package user

import (
	"context"
	_ "embed"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user/avatar"
	"github.com/allocamelus/allocamelus/internal/user/perms"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/jackc/pgx/v5"
)

// User Types
type Types int8

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
	Avatar      string       `msg:"-" json:"avatar"`
	Bio         string       `msg:"bio" json:"bio,omitempty"`
	SelfFollow  FollowStruct `msg:"-" json:"selfFollow,omitempty"`
	UserFollow  FollowStruct `msg:"-" json:"userFollow,omitempty"`
	Followers   int64        `msg:"followers" json:"followers"`
	Type        Types        `msg:"type" json:"type"`
	Permissions perms.Perms  `msg:"permissions" json:"-"`
	Created     int64        `msg:"created" json:"created,omitempty"`
}

// New user
func New(userName, name, email string) *User {
	user := new(User)
	user.UserName = userName
	user.Name = name
	user.Email = email
	user.SetDefaultPerms()
	user.Created = time.Now().Unix()
	return user
}

// Insert new user into database
//
//	returns nil and sets user.ID on success
func (u *User) Insert() error {
	var err error
	// Insert user into database
	u.ID, err = g.Data.Queries.InsertUser(context.Background(), db.InsertUserParams{
		Username:    u.UserName,
		Email:       u.Email,
		Type:        int16(u.Type),
		Permissions: int64(u.Permissions),
		Created:     u.Created,
	})

	return err
}

// GetPublic user info for session user
// TODO: Cache
func GetPublic(s *session.Session, userID int64) (user User, err error) {
	user.ID = userID
	r, err := g.Data.Queries.GetPublicUser(context.Background(), userID)
	if err != nil {
		return
	}
	user.UserName = r.Username
	user.Name = r.Name
	user.Bio = r.Bio
	user.Type = Types(r.Type)
	user.Created = r.Created

	user.Avatar, err = avatar.GetUrl(userID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return
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

// SetDefaultPerms sets default permissions
func (u *User) SetDefaultPerms() {
	u.Permissions = perms.Default
}
