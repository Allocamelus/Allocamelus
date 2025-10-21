package user

import (
	"context"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/g"
)

// GetIDByEmail get user id by email
func GetIDByEmail(email string) (userID int64, err error) {
	return g.Data.Queries.GetUserIDByEmail(context.Background(), email)
}

// GetEmailByID get email by user id
func GetEmailByID(userID int64) (email string, err error) {
	return g.Data.Queries.GetUserEmailByID(context.Background(), userID)
}

// GetUserNameByID get userName by user id
func GetUserNameByID(userID int64) (userName string, err error) {
	return g.Data.Queries.GetUserNameByID(context.Background(), userID)
}

// GetIDByUserName get user id by username
func GetIDByUserName(username string) (userID int64, err error) {
	return g.Data.Queries.GetUserIDByUserName(context.Background(), username)
}
