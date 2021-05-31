//go:generate msgp

package post

import (
	"database/sql"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/user"
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

var ErrNoComment = errors.New("post/comment: Error No Comment Found OR Insufficient permission to view this Comment")

// GetCommentForUser
func GetCommentForUser(commentId int64, u *user.Session) (*Comment, error) {
	c, err := GetComment(commentId)
	if err != nil {
		if err != sql.ErrNoRows {
			return c, err
		}
		return nil, ErrNoComment
	}

	// Is user the commenter
	if compare.EqualInt64(c.UserID, u.UserID) {
		return c, nil
	}

	// Check if user can view post
	if err = CanView(c.PostID, u); err != nil {
		return nil, err
	}

	return c, nil
}

var preGetCommentUserID *sql.Stmt

func GetCommentUserId(commentID int64) (int64, error) {
	if preGetCommentUserID == nil {
		preGetCommentUserID = g.Data.Prepare(`SELECT userId FROM PostComments WHERE postCommentId = ? LIMIT 1`)
	}
	var userId int64
	err := preGetCommentUserID.QueryRow(commentID).Scan(&userId)
	return userId, err
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

var (
	// ErrCommentContentLength max 4096
	ErrCommentContentLength = errors.New("invalid-length-min0-max4096")
)

// Validate is content valid
func (c *Comment) Validate() error {
	return CommentContentValid(c.Content)
}

// CommentContentValid Comment content valid
func CommentContentValid(content string) error {
	if err := validation.Validate(content,
		validation.Length(0, 4096),
	); err != nil {
		return ErrCommentContentLength
	}

	if !g.ContentInvalidChars.MatchString(content) {
		return g.ErrInvalidChars
	}

	return nil
}
