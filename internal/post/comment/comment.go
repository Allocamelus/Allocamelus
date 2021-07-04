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
	// ReplyToId comment id
	ReplyToId int64  `msg:"replyToId" json:"replyToId"`
	Created   int64  `msg:"created" json:"created"`
	Updated   int64  `msg:"updated" json:"updated"`
	Content   string `msg:"content" json:"content"`
	Replies   int64  `msg:"replies" json:"replies"`
}

func New(userID, postID, ReplyToId int64, content string) *Comment {
	comment := new(Comment)
	comment.UserID = userID
	comment.PostID = postID
	comment.ReplyToId = ReplyToId
	comment.Created = time.Now().Unix()
	comment.Content = content
	return comment
}

var preGet *sql.Stmt

// Get
func Get(commentId int64) (*Comment, error) {
	if preGet == nil {
		preGet = g.Data.Prepare(`
	 	SELECT
			postId,
			userId,
			replyToComment,
			created,
			updated,
			content
		FROM PostComments
		WHERE postCommentId = ? LIMIT 1`)
	}
	c := new(Comment)
	c.ID = commentId
	err := preGet.QueryRow(commentId).Scan(&c.PostID, &c.UserID, &c.ReplyToId, &c.Created, &c.Updated, &c.Content)
	if err != nil {
		return nil, err
	}

	// Get reply count if any
	if err := c.CountReplies(); err != nil {
		return nil, err
	}

	return c, err
}

var ErrNoComment = errors.New("post/comment: Error No Comment Found OR Insufficient permission to view this Comment")

// GetForUser
func GetForUser(commentId int64, u *user.Session) (*Comment, error) {
	c, err := Get(commentId)
	if err != nil {
		if err != sql.ErrNoRows {
			return c, err
		}
		return nil, ErrNoComment
	}

	if err = CanView(commentId, u, c); err != nil {
		return nil, err
	}

	return c, nil
}

var preGetUserID *sql.Stmt

func GetUserId(commentID int64) (int64, error) {
	if preGetUserID == nil {
		preGetUserID = g.Data.Prepare(`SELECT userId FROM PostComments WHERE postCommentId = ? LIMIT 1`)
	}
	var userId int64
	err := preGetUserID.QueryRow(commentID).Scan(&userId)
	return userId, err
}

var preGetCanView *sql.Stmt

func CanView(commentID int64, u *user.Session, commentCache ...*Comment) error {
	if preGetCanView == nil {
		preGetCanView = g.Data.Prepare(`SELECT postId, userId FROM PostComments WHERE postCommentId = ? LIMIT 1`)
	}

	var c *Comment
	// Check commentCache
	if len(commentCache) != 0 && commentCache[0] != nil {
		// Use commentCache if valid
		c = commentCache[0]
	} else {
		// Get comment from store
		c = new(Comment)
		c.PostID = commentID
		err := preGetCanView.QueryRow(commentID).Scan(&c.PostID, &c.UserID)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}
			return ErrNoComment
		}
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

func CanReplyTo(commentID int64, u *user.Session, commentCache ...*Comment) error {
	if commentID == 0 {
		return nil
	}
	return CanView(commentID, u, commentCache...)
}

var preInsert *sql.Stmt

func (c *Comment) Insert() error {
	if preInsert == nil {
		preInsert = g.Data.Prepare(`
			INSERT INTO PostComments (
				postId, 
				userId, 
				replyToComment, 
				created, 
				content
			)
			VALUES (?, ?, ?, ?, ?)`)
	}
	r, err := preInsert.Exec(
		c.PostID, c.UserID,
		c.ReplyToId, c.Created,
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
		preCountCommentReplies = g.Data.Prepare(`SELECT COUNT(*) FROM PostComments WHERE replyToComment = ?`)
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
