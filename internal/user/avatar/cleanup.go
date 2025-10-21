package avatar

import (
	"context"
	_ "embed"
	"os"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

func deactivateOld(userId int64) error {
	return g.Data.Queries.DeactivateOldUserAvatars(context.Background(), userId)
}

// CleanupOld removes deactivated avatars from file store by userId
func CleanupOld(userId int64) error {
	rows, err := g.Data.Queries.GetOldUserAvatars(context.Background(), userId)
	if err != nil {
		return err
	}

	for _, r := range rows {
		var (
			fileType fileutil.Format
			b58hash  string
		)

		fileType = fileutil.Format(r.Filetype)
		b58hash = r.Hash

		remove(b58hash, fileType)
	}
	return nil
}

// remove avatar file and db entry
func remove(b58hash string, fileType fileutil.Format) error {
	os.RemoveAll(fileutil.FilePath(selectorPath(b58hash, true)))

	return g.Data.Queries.DeleteUserAvatarByFile(context.Background(), db.DeleteUserAvatarByFileParams{Filetype: int32(fileType), Hash: b58hash})
}
