package user

import (
	"context"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
)

const (
	Unverified Types = iota
	Private
	Public
)

// GetType get user's type
func GetType(userID int64) (t Types, err error) {
	tUn, err := g.Data.Queries.GetUserType(context.Background(), userID)
	t = Types(tUn)
	return
}

// UpdateType update user's type
func UpdateType(userID int64, t Types) error {
	return g.Data.Queries.UpdateUserType(context.Background(), db.UpdateUserTypeParams{Type: int16(t), Userid: userID})
}

// IsVerified is user type != Unverified
func IsVerified(userID int64) (bool, error) {
	t, err := GetType(userID)
	if err != nil {
		return false, err
	}
	return !t.Unverified(), nil
}

// Unverified is type == Unverified
func (t Types) Unverified() bool {
	return (t == Unverified)
}

// Private is type == Private
func (t Types) Private() bool {
	return (t == Private)
}

// Public is type == Public
func (t Types) Public() bool {
	return (t == Public)
}
