package avatar

import (
	"database/sql"
	_ "embed"
	"os"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

var (
	//go:embed sql/deactivateOld.sql
	qDeactivateOld   string
	preDeactivateOld *sql.Stmt
	//go:embed sql/getOld.sql
	qGetOld   string
	preGetOld *sql.Stmt
	//go:embed sql/delete.sql
	qDelete   string
	preDelete *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preDeactivateOld, qDeactivateOld)
	data.PrepareQueuer.Add(&preGetOld, qGetOld)
	data.PrepareQueuer.Add(&preDelete, qDelete)
}

func deactivateOld(userId int64) error {
	_, err := preDeactivateOld.Exec(userId)
	return err
}

// CleanupOld removes deactivated avatars from file store by userId
func CleanupOld(userId int64) error {
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
	os.RemoveAll(fileutil.FilePath(selectorPath(b58hash, true)))

	_, err := preDelete.Exec(fileType, b58hash)
	return err
}
