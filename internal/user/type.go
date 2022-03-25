package user

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
)

const (
	Unverified Types = iota
	Private
	Public
)

var (
	//go:embed sql/get/type.sql
	qGetType   string
	preGetType *sql.Stmt
	//go:embed sql/update/type.sql
	qUpdateType   string
	preUpdateType *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetType, qGetType)
	data.PrepareQueuer.Add(&preUpdateType, qUpdateType)
}

// GetType get user's type
func GetType(userID int64) (t Types, err error) {
	err = preGetType.QueryRow(userID).Scan(&t)
	return
}

// UpdateType update user's type
func UpdateType(userID int64, t Types) error {
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
