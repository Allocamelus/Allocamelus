package avatar

import (
	"database/sql"
	_ "embed"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

var (
	//go:embed sql/hasAvatar.sql
	qHasAvatar   string
	preHasAvatar *sql.Stmt
	//go:embed sql/getAvatarUrl.sql
	qGetAvatarUrl   string
	preGetAvatarUrl *sql.Stmt
	//go:embed sql/insertAvatar.sql
	qInsertAvatar   string
	preInsertAvatar *sql.Stmt
	//go:embed sql/remove.sql
	qRemove   string
	preRemove *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preHasAvatar, qHasAvatar)
	data.PrepareQueuer.Add(&preGetAvatarUrl, qGetAvatarUrl)
	data.PrepareQueuer.Add(&preInsertAvatar, qInsertAvatar)
	data.PrepareQueuer.Add(&preRemove, qRemove)
}

const (
	MaxHightWidth int = 500
	SubPath           = "users/avatars"
)

func HasAvatar(userID int64) (hasAvatar bool, err error) {
	err = preHasAvatar.QueryRow(userID).Scan(&hasAvatar)
	return
}

func GetUrl(userID int64) (url string, err error) {
	var (
		b58hash string
	)
	err = preGetAvatarUrl.QueryRow(userID).Scan(&b58hash)
	if err != nil {
		return
	}
	url = fileutil.PublicPath(selectorPath(b58hash, true))
	return
}

func InsertAvatar(userID int64, fileType fileutil.Format, b58hash string) error {
	_, err := preInsertAvatar.Exec(userID, time.Now().Unix(), fileType, b58hash)
	if err != nil {
		return err
	}

	return nil
}

func Remove(userID int64) error {
	_, err := preRemove.Exec(userID)
	if err != nil {
		return err
	}

	return CleanupOld(userID)
}

func selectorPath(b58hash string, includeFile bool) string {
	return fileutil.RelativePath(SubPath, b58hash, includeFile)
}
