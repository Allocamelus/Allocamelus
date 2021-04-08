package avatar

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	preGetAvatarUrl      *sql.Stmt
	preInsertAvatar      *sql.Stmt
	preCleanupOldAvatars *sql.Stmt
)

const MaxHightWidth uint = 500

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
	if preCleanupOldAvatars == nil {
		preCleanupOldAvatars = g.Data.Prepare(`UPDATE UserAvatars SET active = 0 WHERE userAvatarId != ? AND userID = ? AND active = 1`)
	}
	_, err = preCleanupOldAvatars.Exec(avatarId, userID)
	if err != nil {
		return err
	}
	return err
}
