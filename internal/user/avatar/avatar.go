package avatar

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	preHasAvatar    *sql.Stmt
	preGetAvatarUrl *sql.Stmt
	preInsertAvatar *sql.Stmt
	preRemove       *sql.Stmt
)

const MaxHightWidth uint = 500

func HasAvatar(userID int64) (hasAvatar bool, err error) {
	if preHasAvatar == nil {
		preHasAvatar = g.Data.Prepare(`SELECT EXISTS(SELECT * FROM UserAvatars WHERE userId = ? AND active = 1)`)
	}
	err = preHasAvatar.QueryRow(userID).Scan(&hasAvatar)
	return
}

func GetUrl(userID int64) (url string, err error) {
	if preGetAvatarUrl == nil {
		preGetAvatarUrl = g.Data.Prepare(`SELECT userAvatarId, selector FROM UserAvatars WHERE userId = ? AND active = 1 ORDER BY userAvatarId DESC LIMIT 1`)
	}
	var (
		avatarId int64
		selector string
	)
	err = preGetAvatarUrl.QueryRow(userID).Scan(&avatarId, &selector)
	url = publicPath(avatarId, selector)
	return
}

func InsertAvatar(userID int64, selector string) (int64, error) {
	if preInsertAvatar == nil {
		preInsertAvatar = g.Data.Prepare(`INSERT INTO UserAvatars (userID, created, selector) VALUES (?, ?, ?)`)
	}
	r, err := preInsertAvatar.Exec(userID, time.Now().Unix(), selector)
	if err != nil {
		return 0, err
	}
	avatarId, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return avatarId, nil
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

func selectorPath(avatarId int64, selector string) string {
	return "users/avatars/" + strconv.Itoa(int(avatarId)) + "/" + selector
}

func filePath(avatarId int64, selector string) string {
	return filepath.Join(g.Config.Path.Media, selectorPath(avatarId, selector))
}

func publicPath(avatarId int64, selector string) string {
	return filepath.Join(g.Config.Path.Public.Media, selectorPath(avatarId, selector))
}
