package user

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	//go:embed sql/update/name.sql
	qUpdateName   string
	preUpdateName *sql.Stmt
	//go:embed sql/update/bio.sql
	qUpdateBio   string
	preUpdateBio *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preUpdateName, qUpdateName)
	data.PrepareQueuer.Add(&preUpdateBio, qUpdateBio)
}

func UpdateName(userId int64, name string) error {
	_, err := preUpdateName.Exec(name, userId)
	return err
}

func UpdateBio(userId int64, bio string) error {
	_, err := preUpdateBio.Exec(bio, userId)
	return err
}
