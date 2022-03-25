package user

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	//go:embed sql/idByEmail.sql
	qIDByEmail   string
	preIDByEmail *sql.Stmt
	//go:embed sql/emailByID.sql
	qEmailByID   string
	preEmailByID *sql.Stmt
	//go:embed sql/userNameByID.sql
	qUserNameByID   string
	preUserNameByID *sql.Stmt
	//go:embed sql/idByUserName.sql
	qIDByUserName   string
	preIDByUserName *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preIDByEmail, qIDByEmail)
	data.PrepareQueuer.Add(&preEmailByID, qEmailByID)
	data.PrepareQueuer.Add(&preUserNameByID, qUserNameByID)
	data.PrepareQueuer.Add(&preIDByUserName, qIDByUserName)
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
