//go:generate msgp

// TODO: Markdown Comments

package comment

import (
	"context"
	_ "embed"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/allocamelus/allocamelus/internal/pkg/errtools"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/user/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"k8s.io/klog/v2"
)

type Comment struct {
	ID     int64 `msg:"id" json:"id"`
	UserID int64 `msg:"userId" json:"userId"`
	PostID int64 `msg:"postId" json:"postId"`
	// ParentID comment id
	ParentID int64              `msg:"parentId" json:"parentId"`
	Created  int64              `msg:"created" json:"created"`
	Updated  int64              `msg:"updated" json:"updated"`
	Content  string             `msg:"content" json:"content"`
	Replies  int64              `msg:"replies" json:"replies"`
	Depth    int64              `msg:"depth" json:"depth"`
	Children map[int64]*Comment `msg:"children" json:"children"`
}

type CommentList map[int64]*Comment

type ListComments struct {
	Comments map[int64]*Comment `msg:"comments" json:"comments"`
}

type List struct {
	ListComments
	Order map[int64]int64 `msg:"order" json:"order"`
}

func DBCommentToComment(dbC *db.Postcomment) *Comment {
	return &Comment{
		ID:       dbC.Postcommentid,
		UserID:   dbC.Userid,
		PostID:   dbC.Postid,
		ParentID: dbC.Parent,
		Created:  dbC.Created,
		Updated:  dbC.Updated,
		Content:  dbC.Content,
	}
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
	c.Children = map[int64]*Comment{}
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
	return GetPostUserID(commentID)
}

func CanView(commentID int64, u *session.Session, commentCache ...*Comment) error {
	c, err := canViewCheckCache(commentID, commentCache...)
	if err != nil {
		return err
	}
	// Is user the commenter
	if compare.EqualInt(c.UserID, u.UserID) {
		return nil
	}

	// Check if user can view post
	if err := post.CanView(c.PostID, u); err != nil {
		return err
	}

	return nil
}

func CanReplyTo(commentID, postID int64, u *session.Session, commentCache ...*Comment) error {
	if commentID == 0 {
		return nil
	}
	c, err := canViewCheckCache(commentID, commentCache...)
	if err != nil {
		return err
	}

	if !compare.EqualInt(postID, c.PostID) {
		return ErrNoComment
	}

	return CanView(commentID, u, c)
}

func (c *Comment) Insert() (err error) {
	c.ID, err = g.Data.Queries.InsertPostComment(context.Background(), db.InsertPostCommentParams{
		Postid:  c.PostID,
		Userid:  c.UserID,
		Parent:  c.ParentID,
		Created: c.Created,
		Content: c.Content,
	})
	if err != nil {
		return err
	}

	err = g.Data.Queries.InsertPostCommentClosureSelf(context.Background(), db.InsertPostCommentClosureSelfParams{Parent: c.ID, Child: c.ID})
	if err != nil {
		return err
	}
	if c.ParentID != 0 {
		err = g.Data.Queries.InsertPostCommentClosureDeep(context.Background(), db.InsertPostCommentClosureDeepParams{Parent: c.ParentID, Child: c.ID})
	}
	return
}

func (c *Comment) CountReplies() (err error) {
	if c.ID == 0 {
		klog.Warning("post/comment: comment id == 0 was used")
		return nil
	}
	c.Replies, err = GetRepliesTotal(c.ID)
	return
}

var (
	// ErrContentLength max 4096
	ErrContentLength = errtools.InvalidLen(2, 4096)
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

	if !errtools.ContentInvalidChars.MatchString(content) {
		return errtools.ErrInvalidChars
	}

	return nil
}
