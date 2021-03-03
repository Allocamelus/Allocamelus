package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
)

const (
	// Post perm allows user to post
	Post Perms = 1 << iota
	// UploadMedia perm allows user to upload media
	UploadMedia
	// Admin perm allows user to perform admin actions
	Admin
)

// DefaultPerms default permissions
var DefaultPerms = Post | UploadMedia

// SetDefaultPerms sets default permissions
func (u *User) SetDefaultPerms() {
	u.Permissions = DefaultPerms
}

var (
	preUpdatePerms *sql.Stmt
	preGetPerms    *sql.Stmt
)

func initPerms(p data.Prepare) {
	preUpdatePerms = p(`UPDATE Users SET permissions = ? WHERE userId = ?`)
	preGetPerms = p(`SELECT permissions FROM Users WHERE userId = ? LIMIT 1`)
}

// GetPerms get user's permissions
func GetPerms(userID int64) (perms Perms, err error) {
	err = preGetPerms.QueryRow(userID).Scan(&perms)
	return
}

// UpdatePerms update user's permissions
func UpdatePerms(userID int64, perms Perms) error {
	_, err := preUpdatePerms.Exec(perms, userID)
	return err
}

// IsVerified is user perms != 0
func IsVerified(userID int64) (bool, error) {
	perms, err := GetPerms(userID)
	if err != nil {
		return false, err
	}
	return perms.NotZero(), nil
}

// NotZero is perms != 0
func (p Perms) NotZero() bool {
	return (p != 0)
}

// CanPost can user post
func (p Perms) CanPost() bool {
	return (p&Post == Post)
}

// CanUploadMedia can user upload media
func (p Perms) CanUploadMedia() bool {
	return (p&UploadMedia == UploadMedia)
}

// IsAdmin is user admin
func (p Perms) IsAdmin() bool {
	return (p&Admin == Admin)
}
