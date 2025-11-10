package avatar

import (
	"context"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

const (
	MaxHightWidth int = 500
	SubPath           = "users/avatars"
)

func HasAvatar(userID int64) (hasAvatar bool, err error) {
	return g.Data.Queries.UserHasAvatar(context.Background(), userID)
}

func GetUrl(userID int64) (url string, err error) {
	b58hash, err := g.Data.Queries.GetUserAvatarURLHash(context.Background(), userID)
	if err != nil {
		return
	}

	url = fileutil.PublicPath(selectorPath(b58hash, true))
	return
}

func InsertAvatar(userID int64, fileType fileutil.Format, b58hash string) error {
	return g.Data.Queries.InsertUserAvatar(context.Background(), db.InsertUserAvatarParams{Userid: userID, Created: time.Now().Unix(), Filetype: int32(fileType), Hash: b58hash})
}

func Remove(userID int64) error {
	if err := g.Data.Queries.DeactivateUserAvatar(context.Background(), userID); err != nil {
		return err
	}

	return CleanupOld(userID)
}

func selectorPath(b58hash string, includeFile bool) string {
	return fileutil.RelativePath(SubPath, b58hash, includeFile)
}
