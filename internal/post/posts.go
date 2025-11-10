package post

import (
	"context"
	_ "embed"
	"errors"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/jackc/pgx/v5"
)

// GetPostsTotal
// TODO: Cache!!!
func GetPostsTotal(u *session.Session) (total int64, err error) {
	return g.Data.Queries.CountPublicPostsForUser(context.Background(), u.UserID)
}

// GetPublicPosts
// TODO: Cache
func GetPosts(startNum, perPage int64, u *session.Session) (*List, error) {
	rows, err := g.Data.Queries.GetLatestPublicPostsForUser(context.Background(), db.GetLatestPublicPostsForUserParams{
		Userid: u.UserID,
		Offset: int32(startNum),
		Limit:  int32(perPage),
	})
	if err != nil {
		return nil, err
	}

	posts := NewList()

	for i, r := range rows {
		p := new(Post)
		p.ID = r.Postid
		p.Content = r.Content
		p.UserID = r.Userid
		p.Published = r.Published
		p.Updated = r.Updated

		// Get Media
		p.MediaList, err = media.Get(p.ID)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return nil, err
			}
		}
		p.Media = len(p.MediaList) > 0

		posts.Posts[p.ID] = p
		posts.Order[int64(i)] = p.ID
	}

	return posts, nil
}

// GetUserPostsTotal
// TODO: Cache!!!
func GetUserPostsTotal(userID int64) (int64, error) {
	return g.Data.Queries.CountPublicPostsByUser(context.Background(), userID)
}

func GetUserPosts(userID, startNum, perPage int64) (*List, error) {
	rows, err := g.Data.Queries.GetLatestPublicPostsByUser(context.Background(), db.GetLatestPublicPostsByUserParams{Userid: userID, Offset: int32(startNum), Limit: int32(perPage)})
	if err != nil {
		return nil, err
	}

	posts := NewList()

	for i, r := range rows {
		p := new(Post)
		p.ID = r.Postid
		p.Content = r.Content
		p.UserID = userID
		p.Published = r.Published
		p.Updated = r.Updated

		// Get Media
		p.MediaList, err = media.Get(p.ID)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return nil, err
			}
		}
		p.Media = len(p.MediaList) > 0

		posts.Posts[p.ID] = p
		posts.Order[int64(i)] = p.ID
	}

	return posts, nil
}
