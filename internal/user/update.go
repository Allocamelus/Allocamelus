package user

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/g"
)

var (
	preUpdateName *sql.Stmt
	preUpdateBio  *sql.Stmt
)

func UpdateName(userId int64, name string) error {
	if preUpdateName == nil {
		preUpdateName = g.Data.Prepare(`UPDATE Users SET name = ? WHERE userId = ?`)
	}
	_, err := preUpdateName.Exec(name, userId)
	return err
}

func UpdateBio(userId int64, bio string) error {
	if preUpdateBio == nil {
		preUpdateBio = g.Data.Prepare(`UPDATE Users SET bio = ? WHERE userId = ?`)
	}
	_, err := preUpdateBio.Exec(bio, userId)
	return err
}
