//go:generate msgp

package post

import (
	"context"
	_ "embed"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jackc/pgx/v5"
	"github.com/microcosm-cc/bluemonday"
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
func New(userID int64, content string, publish bool) *Post { // skipcq: RVV-A0005
	p := new(Post)
	p.UserID = userID
	p.Content = content
	p.Created = time.Now().Unix()
	if publish {
		p.Published = time.Now().Unix()
	}
	return p
}

// Insert into database
func (p *Post) Insert() error {
	id, err := g.Data.Queries.InsertPost(context.Background(), db.InsertPostParams{
		Userid:    p.UserID,
		Created:   p.Created,
		Published: p.Published,
		Content:   p.Content,
	})
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

// Get Post
// TODO: Cache
func Get(postID int64) (*Post, error) {
	p := new(Post)
	p.ID = postID
	dbP, err := g.Data.Queries.GetPost(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	p.UserID = dbP.Userid
	p.Created = int64(dbP.Created)
	p.Published = int64(dbP.Published)
	p.Updated = int64(dbP.Updated)
	p.Content = dbP.Content

	// Get Media
	p.MediaList, err = media.Get(postID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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
func GetForUser(postID int64, u *session.Session) (*Post, error) {
	p, err := Get(postID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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

func CanView(postID int64, u *session.Session, postCache ...*Post) error {
	var p *Post
	// Check postCache
	if len(postCache) != 0 && postCache[0] != nil {
		// Use postCache if valid
		p = postCache[0]
	} else {
		// Get post from store
		p = new(Post)
		p.ID = postID
		dbP, err := g.Data.Queries.GetPostCanView(context.Background(), postID)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return err
			}
			return ErrNoPost
		}
		p.UserID = dbP.Userid
		p.Published = int64(dbP.Published)
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
	return g.Data.Queries.GetPostUserID(context.Background(), postID)
}

func Published(postID int64) (bool, error) {
	p, err := g.Data.Queries.GetPostPublishedStatus(context.Background(), postID)
	if err != nil {
		return false, err
	}
	var published bool
	if p > 0 {
		published = true
	}
	return published, nil
}

// Publish post if not already
func (p *Post) Publish() error {
	if !p.IsPublished() {
		return g.Data.Queries.PublishPost(context.Background(), db.PublishPostParams{Published: time.Now().Unix(), Postid: p.ID})
	}
	return nil
}

// MDtoHTMLContent convert markdown to html and sanitize
func (p *Post) MDtoHTMLContent() {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	doc := parser.NewWithExtensions(extensions).Parse([]byte(p.Content))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	renderer := html.NewRenderer(html.RendererOptions{Flags: htmlFlags})

	p.Content = bluemonday.UGCPolicy().Sanitize(
		string(markdown.Render(doc, renderer)),
	)
}

// IsPublished is post draft
func (p *Post) IsPublished() bool {
	return (p.Published > 0)
}

func (p *Post) IsPoster(userID int64) bool {
	return compare.EqualInt(p.UserID, userID)
}
