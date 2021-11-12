package avatar

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

var (
	preHasAvatar    *sql.Stmt
	preGetAvatarUrl *sql.Stmt
	preInsertAvatar *sql.Stmt
	preRemove       *sql.Stmt
)

const (
	MaxHightWidth int = 500
	SubPath           = "users/avatars"
)

func HasAvatar(userID int64) (hasAvatar bool, err error) {
	if preHasAvatar == nil {
		preHasAvatar = g.Data.Prepare(`SELECT EXISTS(SELECT * FROM UserAvatars WHERE userId = ? AND active = 1)`)
	}
	err = preHasAvatar.QueryRow(userID).Scan(&hasAvatar)
	return
}

func GetUrl(userID int64) (url string, err error) {
	if preGetAvatarUrl == nil {
		preGetAvatarUrl = g.Data.Prepare(`SELECT hash FROM UserAvatars WHERE userId = ? AND active = 1 ORDER BY userAvatarId DESC LIMIT 1`)
	}
	var (
		b58hash string
	)
	err = preGetAvatarUrl.QueryRow(userID).Scan(&b58hash)
	if err != nil {
		return
	}
	url = fileutil.PublicPath(selectorPath(b58hash, true))
	return
}

func InsertAvatar(userID int64, fileType fileutil.Format, b58hash string) error {
	if preInsertAvatar == nil {
		preInsertAvatar = g.Data.Prepare(`INSERT INTO UserAvatars (userID, created, fileType, hash) VALUES (?, ?, ?, ?)`)
	}
	_, err := preInsertAvatar.Exec(userID, time.Now().Unix(), fileType, b58hash)
	if err != nil {
		return err
	}

	return nil
}

func Remove(userID int64) error {
	if preRemove == nil {
		preRemove = g.Data.Prepare(`UPDATE UserAvatars SET active = 0 WHERE userID = ? AND active = 1`)
	}
	_, err := preRemove.Exec(userID)
	if err != nil {
		return err
	}

	return CleanupOld(userID)
}

func selectorPath(b58hash string, includeFile bool) string {
	return fileutil.RelativePath(SubPath, b58hash, includeFile)
}
