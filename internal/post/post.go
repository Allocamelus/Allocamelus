//go:generate msgp

package post

import (
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// Post struct
type Post struct {
	ID        int64          `msg:"id" json:"id"`
	UserID    int64          `msg:"userId" json:"userId"`
	Created   int64          `msg:"created" json:"created,omitempty"`
	Published int64          `msg:"published" json:"published"`
	Updated   int64          `msg:"updated" json:"updated"`
	Content   string         `msg:"content" json:"content"`
	Media     bool           `msg:"media" json:"media"`
	MediaList []*media.Media `msg:"mediaList" json:"mediaList,omitempty"`
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
	preInsert       *sql.Stmt
	preGet          *sql.Stmt
	prePublish      *sql.Stmt
	preGetPublished *sql.Stmt
	preGetUserID    *sql.Stmt
)

func initPost(p data.Prepare) {
	preInsert = p(`INSERT INTO Posts (userId, created, published, content)
	VALUES (?, ?, ?, ?)`)
	preGet = p(`SELECT userId, created, published, updated, content FROM Posts WHERE postId = ? LIMIT 1`)
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
func Get(postID int64) (*Post, error) {
	p := new(Post)
	p.ID = postID
	err := preGet.QueryRow(postID).Scan(&p.UserID, &p.Created, &p.Published, &p.Updated, &p.Content)
	if err != nil {
		return nil, err
	}

	// Get Media
	p.MediaList, err = media.Get(postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	p.Media = len(p.MediaList) > 0

	return p, err
}

// Viewing post errors
var (
	ErrNoPost = errors.New("post/post: Error No Post Found OR Insufficient permission to view this post")
)

// GetForUser returns post if user can view it
func GetForUser(postID int64, u *user.Session) (*Post, error) {
	p, err := Get(postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoPost
		}
		return p, err
	}

	if err = CanView(postID, u, p); err != nil {
		return nil, err
	}

	// Omit Created if user is not poster
	if !p.IsPoster(u.UserID) {
		p.Created = 0
	}

	return p, err
}

var preGetCanView *sql.Stmt

func CanView(postID int64, u *user.Session, postCache ...*Post) error {
	if preGetCanView == nil {
		preGetCanView = g.Data.Prepare(`SELECT userId, published FROM Posts WHERE postId = ? LIMIT 1`)
	}

	var p *Post
	// Check postCache
	if len(postCache) != 0 && postCache[0] != nil {
		// Use postCache if valid
		p = postCache[0]
	} else {
		// Get post from store
		p = new(Post)
		p.ID = postID
		err := preGetCanView.QueryRow(postID).Scan(&p.UserID, &p.Published)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}
			return ErrNoPost
		}
	}

	// Check if user can view post
	if !p.IsPublished() {
		if !u.LoggedIn || !p.IsPoster(u.UserID) {
			return user.ErrViewMeNot
		}
	}

	if err := user.CanView(p.UserID, u); err != nil {
		return err
	}

	return nil
}

func GetUserId(postID int64) (int64, error) {
	if preGetUserID == nil {
		preGetUserID = g.Data.Prepare(`SELECT userId FROM Posts WHERE postId = ? LIMIT 1`)
	}
	var userId int64
	err := preGetUserID.QueryRow(postID).Scan(&userId)
	return userId, err
}

func Published(postID int64) (bool, error) {
	if preGetPublished == nil {
		preGetPublished = g.Data.Prepare(`SELECT published FROM Posts WHERE postId = ? LIMIT 1`)
	}

	var published bool
	err := preGetPublished.QueryRow(postID).Scan(&published)
	return published, err
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

func (p *Post) IsPoster(userID int64) bool {
	return compare.EqualInt64(p.UserID, userID)
}
