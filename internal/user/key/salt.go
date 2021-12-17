package key

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
)

func init() {
	data.PrepareQueuer.Add(&preGetSalt, qGetSalt)
}

var (
	//go:embed sql/getSalt.sql
	qGetSalt   string
	preGetSalt *sql.Stmt
)

func GetSalt(userID int64) (salt string, err error) {
	err = preGetSalt.QueryRow(userID).Scan(&salt)
	return
}
