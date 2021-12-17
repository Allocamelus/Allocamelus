//go:generate msgp

package perms

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
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

var (
	//go:embed sql/get.sql
	qGet   string
	preGet *sql.Stmt
	//go:embed sql/update.sql
	qUpdate   string
	preUpdate *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGet, qGet)
	data.PrepareQueuer.Add(&preUpdate, qUpdate)
}

// Get get user's permissions
func Get(userID int64) (perms Perms, err error) {
	err = preGet.QueryRow(userID).Scan(&perms)
	return
}

// Update update user's permissions
func Update(userID int64, perms Perms) error {
	_, err := preUpdate.Exec(perms, userID)
	return err
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
