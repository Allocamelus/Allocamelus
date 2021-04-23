package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/g"
)

const (
	Unverified Types = iota
	Private
	Public
)

var preGetType *sql.Stmt

// GetType get user's type
func GetType(userID int64) (t Types, err error) {
	if preGetType == nil {
		preGetType = g.Data.Prepare(`SELECT type FROM Users WHERE userId = ? LIMIT 1`)
	}
	err = preGetType.QueryRow(userID).Scan(&t)
	return
}

var preUpdateType *sql.Stmt

// UpdateType update user's type
func UpdateType(userID int64, t Types) error {
	if preUpdateType == nil {
		preUpdateType = g.Data.Prepare(`UPDATE Users SET type = ? WHERE userId = ?`)
	}
	_, err := preUpdateType.Exec(t, userID)
	return err
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
