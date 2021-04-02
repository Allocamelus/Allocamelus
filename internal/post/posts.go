package post

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
)

var (
	preGetPublicPosts struct {
		Total  *sql.Stmt
		Latest *sql.Stmt
		ByUser struct {
			Total  *sql.Stmt
			Latest *sql.Stmt
		}
	}
)

func initPosts(p data.Prepare) {
	preGetPublicPosts.Total = p(`SELECT COUNT(postId) FROM Posts WHERE published != 0`)
	preGetPublicPosts.Latest = p(`
	SELECT
		postId, userId, published,
		updated, content, media
	FROM Posts 
	WHERE published != 0 
	ORDER BY published DESC
	LIMIT ?,?`)
	preGetPublicPosts.ByUser.Total = p(`SELECT COUNT(postId) FROM Posts WHERE published != 0 AND userId = ?`)
	preGetPublicPosts.ByUser.Latest = p(`
	SELECT
		postId, published,
		updated, content, media
	FROM Posts 
	WHERE published != 0 AND userId = ?
	ORDER BY published DESC
	LIMIT ?,?`)
}

// GetPublicTotal Posts
// TODO: Cache!!!
func GetPublicTotal() (total int64, err error) {
	err = preGetPublicPosts.Total.QueryRow().Scan(&total)
	return
}

// GetPublicPosts
// TODO: Likes, Views & Cache
func GetPublicPosts(startNum, perPage int64) (*List, error) {
	posts := NewList()

	rows, err := preGetPublicPosts.Latest.Query(startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var index int64
	for rows.Next() {
		p := new(Post)

		err := rows.Scan(&p.ID, &p.UserID, &p.Published, &p.Updated, &p.Content, &p.Media)
		if err != nil {
			return nil, err
		}

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

		err := rows.Scan(&p.ID, &p.Published, &p.Updated, &p.Content, &p.Media)
		if err != nil {
			return nil, err
		}
		p.UserID = userID

		posts.Posts[p.ID] = p
		posts.Order[index] = p.ID
		index++
	}

	return posts, nil
}
