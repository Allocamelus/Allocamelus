package avatar

import (
	"database/sql"
	"os"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

var (
	preDeactivateOld *sql.Stmt
	preGetOld        *sql.Stmt
	preDelete        *sql.Stmt
)

func deactivateOld(userId int64) error {
	if preDeactivateOld == nil {
		preDeactivateOld = g.Data.Prepare(`UPDATE UserAvatars
			SET active = 0
			WHERE userAvatarId IN (
				SELECT userAvatarId FROM (
					SELECT userAvatarId FROM UserAvatars 
					ORDER BY userAvatarId DESC
					LIMIT 0, 18446744073709551615
				) tmp
			) AND userId = ? AND active = 1
			`) // Deactivate all but the latest
	}
	_, err := preDeactivateOld.Exec(userId)
	return err
}

// CleanupOld removes deactivated avatars from file store by userId
func CleanupOld(userId int64) error {
	if preGetOld == nil {
		preGetOld = g.Data.Prepare(`SELECT UA.fileType, UA.hash
			FROM UserAvatars UA
			WHERE userId = ? AND active = 0
				AND NOT EXISTS (
					SELECT *
					FROM UserAvatars
					WHERE UA.hash = hash AND active = 1
				)`)
	}
	rows, err := preGetOld.Query(userId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			fileType fileutil.Format
			b58hash  string
		)
		err := rows.Scan(&fileType, &b58hash)
		if err != nil {
			return err
		}
		remove(b58hash, fileType)
	}
	return nil
}

// remove avatar file and db entry
func remove(b58hash string, fileType fileutil.Format) error {
	os.RemoveAll(fileutil.FilePath(selectorPath(b58hash, fileType, true)))
	if preDelete == nil {
		preDelete = g.Data.Prepare(`DELETE FROM UserAvatars WHERE fileType=? AND hash=?`)
	}
	_, err := preDelete.Exec(fileType, b58hash)
	return err
}
