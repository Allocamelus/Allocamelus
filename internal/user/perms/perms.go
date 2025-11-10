//go:generate msgp

package perms

import (
	"context"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
)

// Perms permissions
type Perms int64

const (
	// Post perm allows user to post
	Post Perms = 1 << iota
	// UploadMedia perm allows user to upload media
	UploadMedia
	// Admin perm allows user to perform admin actions
	Admin
)

// Default default permissions
var Default = Post | UploadMedia

// Get get user's permissions
func Get(userID int64) (Perms, error) {
	p, err := g.Data.Queries.GetUserPermissions(context.Background(), userID)
	if err != nil {
		return 0, err
	}
	return Perms(p), nil
}

// Update update user's permissions
func Update(userID int64, perms Perms) error {
	return g.Data.Queries.UpdateUserPermissions(context.Background(), db.UpdateUserPermissionsParams{Permissions: int64(perms), Userid: userID})
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
