//go:generate msgp
// TODO: Media

package post

import (
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// Post struct
type Post struct {
	ID        int64  `msg:"id" json:"id"`
	UserID    int64  `msg:"userId" json:"userId"`
	Created   int64  `msg:"created" json:"created,omitempty"`
	Published int64  `msg:"published" json:"published"`
	Updated   int64  `msg:"updated" json:"updated"`
	Content   string `msg:"content" json:"content"`
	Media     bool   `msg:"media" json:"media"`
}

// New Post
func New(userID int64, content string, publish bool) *Post {
	p := new(Post)
	p.UserID = userID
	p.Content = content
	p.Created = time.Now().Unix()
	if publish {
		p.Published = time.Now().Unix()
	}
	return p
}

var (
	preInsert         *sql.Stmt
	preGet            *sql.Stmt
	preGetPublicPosts struct {
		Total  *sql.Stmt
		Latest *sql.Stmt
	}
	prePublish *sql.Stmt
)

func initPost(p data.Prepare) {
	preInsert = p(`INSERT INTO Posts (userId, created, published, content)
	VALUES (?, ?, ?, ?)`)
	preGet = p(`SELECT userId, created, published, updated, content, media FROM Posts WHERE postId = ? LIMIT 1`)
	preGetPublicPosts.Total = p(`SELECT COUNT(postId) FROM Posts WHERE published != 0`)
	preGetPublicPosts.Latest = p(`
	SELECT
		postId, userId, published,
		updated, content, media
	FROM Posts 
	WHERE published != 0 
	ORDER BY published DESC
	LIMIT ?,?`)
	prePublish = p(`UPDATE Posts SET published = ? WHERE postId = ?`)
}

// Insert into database
func (p *Post) Insert() error {
	r, err := preInsert.Exec(
		p.UserID, p.Created,
		p.Published, p.Content,
	)
	if err != nil {
		return err
	}

	p.ID, err = r.LastInsertId()
	return err
}

// Get Post
// TODO: Likes, Views & Cache
func Get(postID int64) (Post, error) {
	var p Post
	p.ID = postID
	err := preGet.QueryRow(postID).Scan(&p.UserID, &p.Created, &p.Published, &p.Updated, &p.Content, &p.Media)
	return p, err
}

// GetPublicPosts
// TODO: Likes, Views & Cache
func GetPublicPosts(startNum, perPage int64) ([]*Post, error) {
	var posts []*Post
	rows, err := preGetPublicPosts.Latest.Query(startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := new(Post)

		err := rows.Scan(&p.ID, &p.UserID, &p.Published, &p.Updated, &p.Content, &p.Media)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}

// GetPublicTotal Posts
// TODO: Cache!!!
func GetPublicTotal() (total int64, err error) {
	err = preGetPublicPosts.Total.QueryRow().Scan(&total)
	return
}

// Viewing post errors
var (
	ErrNoPost = errors.New("post/post: Error No Post Found OR Insufficient permission to view this post")
)

// GetForUser returns post if user can view it
func GetForUser(postID int64, u *user.Session) (Post, error) {
	p, err := Get(postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Post{}, ErrNoPost
		}
		return p, err
	}

	// Check if user can view post
	if !p.IsPublished() {
		if !u.LoggedIn || !p.isPoster(u.UserID) {
			return Post{}, ErrNoPost
		}
	}

	// Omit Created if user is not poster
	if !p.isPoster(u.UserID) {
		p.Created = 0
	}

	return p, err
}

// Publish post if not already
func (p *Post) Publish() error {
	if !p.IsPublished() {
		_, err := prePublish.Exec(time.Now().Unix(), p.ID)
		return err
	}
	return nil
}

// MDtoHTMLContent convert markdown to html and sanitize
func (p *Post) MDtoHTMLContent() {
	p.Content = bluemonday.UGCPolicy().Sanitize(
		string(blackfriday.Run([]byte(p.Content))),
	)
}

// IsPublished is post draft
func (p *Post) IsPublished() bool {
	return (p.Published != 0)
}

func (p *Post) isPoster(userID int64) bool {
	return (p.UserID == userID)
}
