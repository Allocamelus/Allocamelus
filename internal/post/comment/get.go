package comment

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user"
)

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

func getForCanView(commentID int64) (*Comment, error) {
	if preGetCanView == nil {
		preGetCanView = g.Data.Prepare(`SELECT postId, userId FROM PostComments WHERE postCommentId = ? LIMIT 1`)
	}

	// Get comment from store
	c := new(Comment)
	c.PostID = commentID
	err := preGetCanView.QueryRow(commentID).Scan(&c.PostID, &c.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, ErrNoComment
	}
	return c, nil
}

var preGetPostTotal *sql.Stmt

// GetTotal comments for post
func GetPostTotal(postID int64) (total int64, err error) {
	if preGetPostTotal == nil {
		preGetPostTotal = g.Data.Prepare(`SELECT COUNT(*) FROM PostComments WHERE postId = ?`)
	}
	err = preGetPostTotal.QueryRow(postID).Scan(&total)
	return
}

const (
	getPostCommentsP1 = `
	SELECT postCommentId, postId, userId, replyToComment, created, updated, content
	FROM PostComments
	WHERE postId = ?
		AND replyToComment IN (
			SELECT postCommentId FROM (
				SELECT postCommentId FROM PostComments`
	getPostCommentsP2 = ` WHERE replyToComment = 0`
	getPostCommentsP3 = `
			) tmp
		)
		OR replyToComment = 0
	LIMIT ?,?`
)

var (
	preGetPostComments     *sql.Stmt
	preGetPostCommentsDeep *sql.Stmt
)

func GetPostComments(startNum, perPage, postID int64, deep bool) (*List, error) {
	if preGetPostComments == nil {
		preGetPostComments = g.Data.Prepare(getPostCommentsP1 + getPostCommentsP2 + getPostCommentsP3)
	}
	if preGetPostCommentsDeep == nil {
		preGetPostCommentsDeep = g.Data.Prepare(getPostCommentsP1 + getPostCommentsP3)
	}

	var query *sql.Stmt

	if deep {
		query = preGetPostCommentsDeep
	} else {
		query = preGetPostComments
	}

	rows, err := query.Query(postID, startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := NewList()
	var i int64
	for rows.Next() {
		c := new(Comment)
		err = rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.ReplyToId, &c.Created, &c.Updated, &c.Content)
		if err != nil {
			return nil, err
		}

		if err = c.CountReplies(); err != nil {
			return nil, err
		}

		l.Comments[c.ID] = c
		l.Order[i] = c.ID
		i++
	}

	return l, nil
}

const (
	getRepliesP1 = `
	SELECT postCommentId, postId, userId, replyToComment, created, updated, content
	FROM PostComments
	WHERE postId = ?
		AND replyToComment IN (
			SELECT postCommentId FROM (
				SELECT postCommentId FROM PostComments`
	getRepliesP2 = ` WHERE replyToComment = 0`
	getRepliesP3 = `
			) tmp
		)
		OR replyToComment = 0
	LIMIT ?,?`
)

var (
	preGetReplies     *sql.Stmt
	preGetRepliesDeep *sql.Stmt
)

func GetReplies(startNum, perPage, postID int64, deep bool) (*List, error) {
	return nil, nil
}
