package comment

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/g"
)

var preDelete *sql.Stmt

// Delete commentID
func Delete(commentID int64) error {
	if preDelete == nil {
		preDelete = g.Data.Prepare(`DELETE FROM PostComments WHERE postCommentId = ?`)
	}

	_, err := preDelete.Exec(commentID)
	return err
}
