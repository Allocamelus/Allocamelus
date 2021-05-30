//go:generate msgp

package post

import (
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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
}

func NewComment(userID, postID, ReplyToId int64, content string) *Comment {
	comment := new(Comment)
	comment.UserID = userID
	comment.PostID = postID
	comment.ReplyToId = ReplyToId
	comment.Created = time.Now().Unix()
	comment.Content = content
	return comment
}

var preGetComment *sql.Stmt

// GetComment
func GetComment(commentId int64) (*Comment, error) {
	if preGetComment == nil {
		preGetComment = g.Data.Prepare(`
	 	SELECT
			postId,
			userId,
			replyToComment,
			created,
			content
		FROM PostComments
		WHERE postCommentId = ? LIMIT 1`)
	}
	c := new(Comment)
	c.ID = commentId
	err := preGet.QueryRow(commentId).Scan(&c.PostID, &c.UserID, &c.ReplyToId, &c.Created, &c.Content)
	return c, err
}

var preInsertComment *sql.Stmt

func (c *Comment) Insert() error {
	if preInsertComment == nil {
		preInsertComment = g.Data.Prepare(`
			INSERT INTO PostComments (
				postId, 
				userId, 
				replyToComment, 
				created, 
				content
			)
			VALUES (?, ?, ?, ?)`)
	}
	r, err := preInsertComment.Exec(
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

// ErrCommentContentLength max 4096
var ErrCommentContentLength = errors.New("invalid-length-min0-max4096")

// Validate is content valid
func (c *Comment) Validate() error {
	return CommentContentValid(c.Content)
}

// CommentContentValid Comment content valid
func CommentContentValid(content string) error {
	if err := validation.Validate(content,
		validation.Length(0, 4096),
	); err != nil {
		return ErrContentLength
	}
	return nil
}
