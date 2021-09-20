package comment

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	//go:embed sql/delete.sql
	qDelete   string
	preDelete *sql.Stmt
)

// Delete commentID
func Delete(commentID int64) error {
	if preDelete == nil {
		preDelete = g.Data.Prepare(qDelete)
	}

	_, err := preDelete.Exec(commentID)
	return err
}
