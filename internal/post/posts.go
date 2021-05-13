package post

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user"
)

var (
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

func initPosts(p data.Prepare) {
	preGetPublicPosts.Total = p(`
	SELECT COUNT(postId)
	FROM Posts
	WHERE userId IN (
		SELECT userId FROM (
			SELECT userId FROM Users 
			WHERE type = 2
		) tmp
	) AND published != 0`)
	preGetPublicPosts.Latest = p(`
	SELECT
		postId, userId, published,
		updated, content
	FROM Posts 
	WHERE userId IN (
		SELECT userId FROM (
			SELECT userId FROM Users 
			WHERE type = 2
		) tmp
	) AND published != 0
	ORDER BY published DESC
	LIMIT ?,?`)

	preGetPublicPosts.ForUser.Total = p(`
	SELECT COUNT(postId)
	FROM Posts
	WHERE userId IN (
		SELECT followUserId FROM (
			SELECT followUserId FROM UserFollows 
			WHERE userId = ? AND accepted = 1
		) tmp
	)
	AND published != 0
	OR userId = ?`)
	preGetPublicPosts.ForUser.Latest = p(`
	SELECT
		postId, userId, published,
		updated, content
	FROM Posts
	WHERE userId IN (
		SELECT followUserId FROM (
			SELECT followUserId FROM UserFollows 
			WHERE userId = ? AND accepted = 1
		) tmp
	)
	AND published != 0
	OR userId = ?
	ORDER BY published DESC
	LIMIT ?,?`)

	preGetPublicPosts.ByUser.Total = p(`SELECT COUNT(postId) FROM Posts WHERE published != 0 AND userId = ?`)
	preGetPublicPosts.ByUser.Latest = p(`
	SELECT
		postId, published,
		updated, content
	FROM Posts 
	WHERE published != 0 AND userId = ?
	ORDER BY published DESC
	LIMIT ?,?`)
}

// GetPublicTotal Posts
// TODO: Cache!!!
func GetPublicTotal(u *user.Session) (total int64, err error) {
	if !u.LoggedIn {
		err = preGetPublicPosts.Total.QueryRow().Scan(&total)
	} else {
		err = preGetPublicPosts.ForUser.Total.QueryRow(u.UserID, u.UserID).Scan(&total)
	}
	return
}

// GetPublicPosts
// TODO: Likes, Views & Cache
func GetPublicPosts(startNum, perPage int64, u *user.Session) (*List, error) {
	posts := NewList()
	var (
		rows *sql.Rows
		err  error
	)

	if !u.LoggedIn {
		rows, err = preGetPublicPosts.Latest.Query(startNum, perPage)
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
