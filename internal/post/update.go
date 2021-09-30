package post

import (
	"database/sql"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	//go:embed sql/updateContent.sql
	qUpdateContent   string
	preUpdateContent *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preUpdateContent, qUpdateContent)
}

func UpdateContent(postID int64, content string) error {
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
