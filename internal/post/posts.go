package post

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user"
)

var (
	//go:embed sql/getPublicPosts/Total.sql
	qGetPubPostsTotal string
	//go:embed sql/getPublicPosts/Latest.sql
	qGetPubPostsLatest string
	//go:embed sql/getPublicPosts/ByUser/Total.sql
	qGetPubPostsByUserTotal string
	//go:embed sql/getPublicPosts/ByUser/Latest.sql
	qGetPubPostsByUserLatest string
	//go:embed sql/getPublicPosts/ForUser/Total.sql
	qGetPubPostsForUserTotal string
	//go:embed sql/getPublicPosts/ForUser/Latest.sql
	qGetPubPostsForUserLatest string

	preGetPublicPosts struct {
		Total  *sql.Stmt
		Latest *sql.Stmt
		ByUser struct {
			Total  *sql.Stmt
			Latest *sql.Stmt
		}
		ForUser struct {
			Total  *sql.Stmt
			Latest *sql.Stmt
		}
	}
)

func init() {
	data.PrepareQueuer.Add(&preGetPublicPosts.Total, qGetPubPostsTotal)
	data.PrepareQueuer.Add(&preGetPublicPosts.Latest, qGetPubPostsLatest)
	data.PrepareQueuer.Add(&preGetPublicPosts.ByUser.Total, qGetPubPostsByUserTotal)
	data.PrepareQueuer.Add(&preGetPublicPosts.ByUser.Latest, qGetPubPostsByUserLatest)
	data.PrepareQueuer.Add(&preGetPublicPosts.ForUser.Total, qGetPubPostsForUserTotal)
	data.PrepareQueuer.Add(&preGetPublicPosts.ForUser.Latest, qGetPubPostsForUserLatest)
}

// GetPublicTotal Posts
// TODO: Cache!!!
func GetPublicTotal(u *user.Session) (total int64, err error) {
	if !u.LoggedIn {
		err = preGetPublicPosts.Total.QueryRow(user.Public).Scan(&total)
	} else {
		err = preGetPublicPosts.ForUser.Total.QueryRow(u.UserID, u.UserID).Scan(&total)
	}
	return
}

// GetPublicPosts
// TODO: Cache
func GetPublicPosts(startNum, perPage int64, u *user.Session) (*List, error) {
	posts := NewList()
	var (
		rows *sql.Rows
		err  error
	)

	if !u.LoggedIn {
		rows, err = preGetPublicPosts.Latest.Query(user.Public, startNum, perPage)
	} else {
		rows, err = preGetPublicPosts.ForUser.Latest.Query(u.UserID, u.UserID, startNum, perPage)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var index int64
	for rows.Next() {
		p := new(Post)

		err := rows.Scan(&p.ID, &p.UserID, &p.Published, &p.Updated, &p.Content)
		if err != nil {
			return nil, err
		}

		// Get Media
		p.MediaList, err = media.Get(p.ID)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
		p.Media = len(p.MediaList) > 0

		posts.Posts[p.ID] = p
		posts.Order[index] = p.ID
		index++
	}

	return posts, nil
}

// GetPublicUserTotal Posts
// TODO: Cache!!!
func GetPublicUserTotal(userID int64) (total int64, err error) {
	err = preGetPublicPosts.ByUser.Total.QueryRow(userID).Scan(&total)
	return
}

func GetPublicUserPosts(userID, startNum, perPage int64) (*List, error) {
	posts := NewList()

	rows, err := preGetPublicPosts.ByUser.Latest.Query(userID, startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var index int64
	for rows.Next() {
		p := new(Post)

		err := rows.Scan(&p.ID, &p.Published, &p.Updated, &p.Content)
		if err != nil {
			return nil, err
		}
		p.UserID = userID

		// Get Media
		p.MediaList, err = media.Get(p.ID)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
		p.Media = len(p.MediaList) > 0

		posts.Posts[p.ID] = p
		posts.Order[index] = p.ID
		index++
	}

	return posts, nil
}
