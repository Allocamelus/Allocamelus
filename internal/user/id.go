package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	preIDByEmail    *sql.Stmt
	preEmailByID    *sql.Stmt
	preUserNameByID *sql.Stmt
	preIDByUserName *sql.Stmt
)

func initID(p data.Prepare) {
	preIDByEmail = p(`SELECT userId FROM Users WHERE email = ? LIMIT 1`)
	preEmailByID = p(`SELECT email FROM Users WHERE userId = ? LIMIT 1`)
	preUserNameByID = p(`SELECT userName FROM Users WHERE userId = ? LIMIT 1`)
	preIDByUserName = p(`SELECT userId FROM Users WHERE userName = ? LIMIT 1`)
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

// GetUserNameByID get userName by user id
func GetUserNameByID(userID int64) (userName string, err error) {
	err = preUserNameByID.QueryRow(userID).Scan(&userName)
	return
}

// GetIDByUserName get user id by username
func GetIDByUserName(username string) (userID int64, err error) {
	err = preIDByUserName.QueryRow(username).Scan(&userID)
	return
}
