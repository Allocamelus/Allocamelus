package comment

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	preUpdateContent *sql.Stmt
)

func UpdateContent(commentID int64, content string) error {
	if preUpdateContent == nil {
		preUpdateContent = g.Data.Prepare(`UPDATE PostComments SET updated = ?, content = ? WHERE postCommentId = ?`)
	}

	_, err := preUpdateContent.Exec(time.Now().Unix(), content, commentID)
	return err
}
