package comment

import (
	"context"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
)

func UpdateContent(commentID int64, content string) error {
	return g.Data.Queries.UpdatePostCommentContent(context.Background(), db.UpdatePostCommentContentParams{Updated: time.Now().Unix(), Content: content, Postcommentid: commentID})
}
