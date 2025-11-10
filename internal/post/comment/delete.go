package comment

import (
	"context"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/g"
)

// Delete commentID
func Delete(commentID int64) error {
	return g.Data.Queries.DeletePostComment(context.Background(), commentID)
}
