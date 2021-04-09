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
		preGetAvatarUrl = g.Data.Prepare(`SELECT location FROM UserAvatars WHERE userId = ? AND active = 1 ORDER BY userAvatarId DESC LIMIT 1`)
	}
	var urlPart string
	err = preGetAvatarUrl.QueryRow(userID).Scan(&urlPart)
	url = g.Config.Path.Public.Media + urlPart
	return
}

func InsertAvatar(userID int64, location string) error {
	if preInsertAvatar == nil {
		preInsertAvatar = g.Data.Prepare(`INSERT INTO UserAvatars (userID, created, location) VALUES (?, ?, ?)`)
	}
	r, err := preInsertAvatar.Exec(userID, time.Now().Unix(), location)
	if err != nil {
		return err
	}
	avatarId, err := r.LastInsertId()
	if err != nil {
		return err
	}

	return deactivateOld(userID, avatarId)
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

const pathPart = "users/avatars/"

func locationPath(userId int64, filename string) string {
	return pathPart + strconv.Itoa(int(userId)) + "/" + filename
}

func filePath(location string) string {
	return filepath.Join(g.Config.Path.Media, location)
}
