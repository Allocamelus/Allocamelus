package user

import (
	"context"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
)

func UpdateName(userId int64, name string) error {
	return g.Data.Queries.UpdateUserName(context.Background(), db.UpdateUserNameParams{Name: name, Userid: userId})
}

func UpdateBio(userId int64, bio string) error {
	return g.Data.Queries.UpdateUserBio(context.Background(), db.UpdateUserBioParams{Bio: bio, Userid: userId})
}
