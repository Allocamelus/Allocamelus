package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	preIDByEmail      *sql.Stmt
	preEmailByID      *sql.Stmt
	preUniqueNameByID *sql.Stmt
	preIDByUniqueName *sql.Stmt
)

func initID(p data.Prepare) {
	preIDByEmail = p(`SELECT userId FROM Users WHERE email = ? LIMIT 1`)
	preEmailByID = p(`SELECT email FROM Users WHERE userId = ? LIMIT 1`)
	preUniqueNameByID = p(`SELECT uniqueName FROM Users WHERE userId = ? LIMIT 1`)
	preIDByUniqueName = p(`SELECT userId FROM Users WHERE uniqueName = ? LIMIT 1`)
}

// GetIDByEmail get user id by email
func GetIDByEmail(email string) (userID int64, err error) {
	err = preIDByEmail.QueryRow(email).Scan(&userID)
	return
}

// GetEmailByID get email by user id
func GetEmailByID(userID int64) (email string, err error) {
	err = preEmailByID.QueryRow(userID).Scan(&email)
	return
}

// GetUniqueNameByID get uniqueName by user id
func GetUniqueNameByID(userID int64) (uniqueName string, err error) {
	err = preUniqueNameByID.QueryRow(userID).Scan(&uniqueName)
	return
}

// GetIDByUniqueName get user id by uniquename
func GetIDByUniqueName(uniquename string) (userID int64, err error) {
	err = preIDByUniqueName.QueryRow(uniquename).Scan(&userID)
	return
}
