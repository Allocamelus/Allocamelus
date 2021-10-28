package post

import (
	"database/sql"
	_ "embed"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user"
)

var (
	//go:embed sql/get/publicPosts/ByUser/Total.sql
	qGetPubPostsByUserTotal string
	//go:embed sql/get/publicPosts/ByUser/Latest.sql
	qGetPubPostsByUserLatest string
	//go:embed sql/get/publicPosts/ForUser/Total.sql
	qGetPubPostsForUserTotal string
	//go:embed sql/get/publicPosts/ForUser/Latest.sql
	qGetPubPostsForUserLatest string

	preGetPosts struct {
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
	data.PrepareQueuer.Add(&preGetPosts.ByUser.Total, qGetPubPostsByUserTotal)
	data.PrepareQueuer.Add(&preGetPosts.ByUser.Latest, qGetPubPostsByUserLatest)
	data.PrepareQueuer.Add(&preGetPosts.ForUser.Total, qGetPubPostsForUserTotal)
	data.PrepareQueuer.Add(&preGetPosts.ForUser.Latest, qGetPubPostsForUserLatest)
}

// GetPostsTotal
// TODO: Cache!!!
func GetPostsTotal(u *user.Session) (total int64, err error) {
	err = preGetPosts.ForUser.Total.QueryRow(u.UserID, u.UserID).Scan(&total)
	return
}

// GetPublicPosts
// TODO: Cache
func GetPosts(startNum, perPage int64, u *user.Session) (*List, error) {
	rows, err := preGetPosts.ForUser.Latest.Query(u.UserID, u.UserID, startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := NewList()

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

// GetUserPostsTotal
// TODO: Cache!!!
func GetUserPostsTotal(userID int64) (total int64, err error) {
	err = preGetPosts.ByUser.Total.QueryRow(userID).Scan(&total)
	return
}

func GetUserPosts(userID, startNum, perPage int64) (*List, error) {
	posts := NewList()

	rows, err := preGetPosts.ByUser.Latest.Query(userID, startNum, perPage)
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
