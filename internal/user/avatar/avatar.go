package avatar

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/g"
)

var preGetAvatarUrl *sql.Stmt

func GetUrl(userID int64) (url string, err error) {
	if preGetAvatarUrl == nil {
		preGetAvatarUrl = g.Data.Prepare(`SELECT location FROM UserAvatars WHERE userId = ? ORDER BY userAvatarId DESC LIMIT 1`)
	}
	var urlPart string
	err = preGetAvatarUrl.QueryRow(userID).Scan(&urlPart)
	url = g.Config.Path.Public.Media + urlPart
	return
}

func NewAvatar(userID int64, file string) {

}
