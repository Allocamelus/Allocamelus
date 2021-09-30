package comment

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	//go:embed sql/delete.sql
	qDelete   string
	preDelete *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preDelete, qDelete)
}

// Delete commentID
func Delete(commentID int64) error {
	_, err := preDelete.Exec(commentID)
	return err
}
