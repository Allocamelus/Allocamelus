package user

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
