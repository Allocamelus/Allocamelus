package comment

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

func UpdateContent(commentID int64, content string) error {
	_, err := preUpdateContent.Exec(time.Now().Unix(), content, commentID)
	return err
}
