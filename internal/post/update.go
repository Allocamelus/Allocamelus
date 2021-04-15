package post

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	preUpdateContent *sql.Stmt
)

func UpdateContent(postID int64, content string) error {
	if preUpdateContent == nil {
		preUpdateContent = g.Data.Prepare(`UPDATE Posts SET updated = ?, content = ? WHERE postId = ?`)
	}
	published, err := Published(postID)
	if err != nil {
		return err
	}

	var updated int64
	if published {
		updated = time.Now().Unix()
	}
	_, err = preUpdateContent.Exec(updated, content, postID)
	return err
}
