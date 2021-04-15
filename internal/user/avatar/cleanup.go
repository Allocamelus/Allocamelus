package avatar

import (
	"database/sql"
	"os"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
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
			preGetOld = g.Data.Prepare(`SELECT userAvatarId FROM UserAvatars WHERE userId = ? AND active = 0`)
		}
		rows, err := preGetOld.Query(userId)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var avatarId int64
			err := rows.Scan(&avatarId)
			if err != nil {
				return err
			}
			remove(avatarId)
		}

	}
	return nil
}

// remove avatar file and db entry
func remove(avatarId int64) error {
	os.RemoveAll(fileutil.FilePath(selectorPath(avatarId, "")))
	if preDelete == nil {
		preDelete = g.Data.Prepare(`DELETE FROM UserAvatars WHERE userAvatarId=?`)
	}
	_, err := preDelete.Exec(avatarId)
	return err
}
