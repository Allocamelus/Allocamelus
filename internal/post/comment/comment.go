//go:generate msgp

package comment

import (
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"k8s.io/klog/v2"
)

type Comment struct {
	ID     int64 `msg:"id" json:"id"`
	UserID int64 `msg:"userId" json:"userId"`
	PostID int64 `msg:"postId" json:"postId"`
	// ParentID comment id
	ParentID int64           `msg:"parentId" json:"parentId"`
	Created  int64           `msg:"created" json:"created"`
	Updated  int64           `msg:"updated" json:"updated"`
	Content  string          `msg:"content" json:"content"`
	Replies  int64           `msg:"replies" json:"replies"`
	Children map[int64]int64 `msg:"children" json:"children"`
}

type ListComments struct {
	Comments map[int64]*Comment `msg:"comments" json:"comments"`
}

type List struct {
	ListComments
	Order map[int64]int64 `msg:"order" json:"order"`
}

func New(userID, postID, ParentID int64, content string) *Comment {
	comment := newComment()
	comment.UserID = userID
	comment.PostID = postID
	comment.ParentID = ParentID
	comment.Created = time.Now().Unix()
	comment.Content = content
	return comment
}

func newComment() *Comment {
	c := new(Comment)
	c.Children = map[int64]int64{}
	return c
}

func NewList() *List {
	l := new(List)
	l.Comments = map[int64]*Comment{}
	l.Order = map[int64]int64{}
	return l
}

var ErrNoComment = errors.New("post/comment: Error No Comment Found OR Insufficient permission to view this Comment")

func canViewCheckCache(commentID int64, commentCache ...*Comment) (*Comment, error) {
	// Check commentCache
	if len(commentCache) != 0 && commentCache[0] != nil {
		// Use commentCache if valid
		return commentCache[0], nil
	}
	return getForCanView(commentID)
}

func CanView(commentID int64, u *user.Session, commentCache ...*Comment) error {
	c, err := canViewCheckCache(commentID, commentCache...)
	if err != nil {
		return err
	}
	// Is user the commenter
	if compare.EqualInt64(c.UserID, u.UserID) {
		return nil
	}

	// Check if user can view post
	if err := post.CanView(c.PostID, u); err != nil {
		return err
	}

	return nil
}

func CanReplyTo(commentID, postID int64, u *user.Session, commentCache ...*Comment) error {
	if commentID == 0 {
		return nil
	}
	c, err := canViewCheckCache(commentID, commentCache...)
	if err != nil {
		return err
	}

	if !compare.EqualInt64(postID, c.PostID) {
		return ErrNoComment
	}

	return CanView(commentID, u, c)
}

var preInsert *sql.Stmt

func (c *Comment) Insert() error {
	if preInsert == nil {
		preInsert = g.Data.Prepare(`
			INSERT INTO PostComments (
				postId, 
				userId, 
				parentComment, 
				created, 
				content
			)
			VALUES (?, ?, ?, ?, ?)`)
	}
	r, err := preInsert.Exec(
		c.PostID, c.UserID,
		c.ParentID, c.Created,
		c.Content,
	)
	if err != nil {
		return err
	}

	c.ID, err = r.LastInsertId()
	return err
}

var preCountCommentReplies *sql.Stmt

func (c *Comment) CountReplies() error {
	if preCountCommentReplies == nil {
		preCountCommentReplies = g.Data.Prepare(`SELECT COUNT(*) FROM PostComments WHERE parentComment = ?`)
	}
	if c.ID == 0 {
		if klog.V(4).Enabled() {
			klog.Info("post/comment: comment id == 0 was used")
		}
		return nil
	}
	err := preCountCommentReplies.QueryRow(c.ID).Scan(&c.Replies)
	return err
}

var (
	// ErrContentLength max 4096
	ErrContentLength = errors.New("invalid-length-min2-max4096")
)

// Validate is content valid
func (c *Comment) Validate() error {
	return ContentValid(c.Content)
}

// ContentValid Comment content valid
func ContentValid(content string) error {
	if err := validation.Validate(content,
		validation.Required,
		validation.Length(2, 4096),
	); err != nil {
		return ErrContentLength
	}

	if !g.ContentInvalidChars.MatchString(content) {
		return g.ErrInvalidChars
	}

	return nil
}
