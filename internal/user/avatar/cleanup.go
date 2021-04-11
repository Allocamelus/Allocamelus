package avatar

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	preDeactivateOld *sql.Stmt
	preHasOld        *sql.Stmt
	preGetOld        *sql.Stmt
	preDelete        *sql.Stmt
)

func deactivateOld(userId, currentAvatarId int64) error {
	if preDeactivateOld == nil {
		preDeactivateOld = g.Data.Prepare(`UPDATE UserAvatars SET active = 0 WHERE userAvatarId != ? AND userID = ? AND active = 1`)
	}
	_, err := preDeactivateOld.Exec(currentAvatarId, userId)
	return err
}

// CleanupOld removes deactivated avatars from file store by userId
func CleanupOld(userId int64) error {
	if preHasOld == nil {
		preHasOld = g.Data.Prepare(`SELECT EXISTS(SELECT * FROM UserAvatars WHERE userId = ? AND active = 0)`)
	}
	var hasOld bool
	err := preHasOld.QueryRow(userId).Scan(&hasOld)
	if err != nil {
		return err
	}
	if hasOld {
		if preGetOld == nil {
			preGetOld = g.Data.Prepare(`SELECT userAvatarId, location FROM UserAvatars WHERE userId = ? AND active = 0`)
		}
		rows, err := preGetOld.Query(userId)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				avatarId int64
				location string
			)
			err := rows.Scan(&avatarId, &location)
			if err != nil {
				return err
			}
			remove(avatarId, userId, location)
		}

	}
	return nil
}

// remove avatar file and db entry
func remove(avatarId, userId int64, location string) error {
	match, err := filepath.Match(locationPath(userId, "*"), location)
	println(err)
	if match {
		os.RemoveAll(filePath(location))
	}
	if preDelete == nil {
		preDelete = g.Data.Prepare(`DELETE FROM UserAvatars WHERE userAvatarId=?`)
	}
	_, err = preDelete.Exec(avatarId)
	return err
}