package key

import (
	"context"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/g"
)

func GetSalt(userID int64) (string, error) {
	return g.Data.Queries.GetUserAuthKeySalt(context.Background(), userID)
}
