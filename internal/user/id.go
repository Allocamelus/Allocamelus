package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	preIDByEmail      *sql.Stmt
	preIDByUniqueName *sql.Stmt
)

func initID(p data.Prepare) {
	preIDByEmail = p(`SELECT userId FROM Users WHERE email = ? LIMIT 1`)
	preIDByUniqueName = p(`SELECT userId FROM Users WHERE uniqueName = ? LIMIT 1`)
}

// GetIDByEmail get user id by email
func GetIDByEmail(email string) (userID int64, err error) {
	err = preIDByEmail.QueryRow(email).Scan(&userID)
	return
}

// GetIDByUniqueName get user id by uniquename
func GetIDByUniqueName(uniquename string) (userID int64, err error) {
	err = preIDByUniqueName.QueryRow(uniquename).Scan(&userID)
	return
}
