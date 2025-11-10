package post

import (
	"context"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
)

func UpdateContent(postID int64, content string) error {
	published, err := Published(postID)
	if err != nil {
		return err
	}

	var updated int64
	if published {
		updated = time.Now().Unix()
	}

	return g.Data.Queries.UpdatePostContent(context.Background(), db.UpdatePostContentParams{
		Updated: updated,
		Content: content,
		Postid:  postID,
	})
}
