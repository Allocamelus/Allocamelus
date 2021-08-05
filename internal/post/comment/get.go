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
			parent,
			created,
			updated,
			content
		FROM PostComments
		WHERE postCommentId = ? LIMIT 1`)
	}
	c := newComment()
	c.ID = commentId
	err := preGet.QueryRow(commentId).Scan(&c.PostID, &c.UserID, &c.ParentID, &c.Created, &c.Updated, &c.Content)
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
	c := newComment()
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

const (
	// Query parts
	// Count Comments from PostComments
	countPostComments = `
	SELECT COUNT(*)
	FROM PostComments`
	countClosures = `
	SELECT COUNT(*)
	FROM PostCommentClosures`
	// Select All Columns from PostComments
	selectPostComments = `
	SELECT postCommentId, postId, userId, parent, created, updated, content
	FROM PostComments`

	queryOrderLimit = ` ORDER BY postCommentId ASC LIMIT ?,?`

	// Get comments and replies
	getPostCommentsP1 = ` WHERE postId = ?`

	// Build queries
	// Post query parts
	partGetPostComments = getPostCommentsP1 + `
	AND parent IN (
		SELECT postCommentId FROM (
			SELECT postCommentId FROM PostComments
			WHERE parent = 0
			) tmp
		)
	OR parent = 0`

	// Get Post queries
	queryGetPostComments     = selectPostComments + partGetPostComments + queryOrderLimit
	queryGetPostCommentsDeep = selectPostComments + getPostCommentsP1 + queryOrderLimit

	// Total Post queries
	queryGetPostTotal = countPostComments + partGetPostComments
)

var preGetPostTotal *sql.Stmt

// GetTotal comments for post
func GetPostTotal(postID int64) (total int64, err error) {
	if preGetPostTotal == nil {
		preGetPostTotal = g.Data.Prepare(queryGetPostTotal)
	}

	err = preGetPostTotal.QueryRow(postID).Scan(&total)
	return
}

var (
	preGetPostComments     *sql.Stmt
	preGetPostCommentsDeep *sql.Stmt
)

func GetPostComments(startNum, perPage, postID int64, deep bool) (*List, error) {
	if preGetPostComments == nil {
		preGetPostComments = g.Data.Prepare(queryGetPostComments)
	}
	if preGetPostCommentsDeep == nil {
		preGetPostCommentsDeep = g.Data.Prepare(queryGetPostCommentsDeep)
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

	return iterRowToList(rows)
}

const (
	getRepliesP1 = `
	WHERE parent = ?`
	getRepliesP2 = `
		OR parent IN (
			SELECT postCommentId FROM (
				SELECT postCommentId FROM PostComments
	 			WHERE parent = ?
			) tmp
		)`

	// Build queries
	// Reply query parts
	partGetRepliesComments     = getRepliesP1
	partGetRepliesCommentsDeep = getRepliesP1 + getRepliesP2

	// Total Replies queries
	queryGetRepliesTotal     = countPostComments + partGetRepliesComments
	queryGetRepliesTotalDeep = countPostComments + partGetRepliesCommentsDeep

	// Get Replies queries
	queryGetReplies     = selectPostComments + partGetRepliesComments + queryOrderLimit
	queryGetRepliesDeep = selectPostComments + partGetRepliesCommentsDeep + queryOrderLimit
)

var (
	preGetReplies          *sql.Stmt
	preGetRepliesDeep      *sql.Stmt
	preGetRepliesTotal     *sql.Stmt
	preGetRepliesTotalDeep *sql.Stmt
)

// GetRepliesTotal
func GetRepliesTotal(commentID int64, deep bool) (total int64, err error) {
	if preGetRepliesTotal == nil {
		preGetRepliesTotal = g.Data.Prepare(queryGetRepliesTotal)
	}
	if preGetRepliesTotalDeep == nil {
		preGetRepliesTotalDeep = g.Data.Prepare(queryGetRepliesTotalDeep)
	}

	var row *sql.Row

	if deep {
		row = preGetRepliesTotalDeep.QueryRow(commentID, commentID)
	} else {
		row = preGetRepliesTotal.QueryRow(commentID)
	}

	err = row.Scan(&total)
	return
}

func GetReplies(startNum, perPage, commentID int64, deep bool) (*List, error) {
	if preGetReplies == nil {
		preGetReplies = g.Data.Prepare(queryGetReplies)
	}
	if preGetRepliesDeep == nil {
		preGetRepliesDeep = g.Data.Prepare(queryGetRepliesDeep)
	}

	var (
		rows *sql.Rows
		err  error
	)

	if deep {
		rows, err = preGetRepliesDeep.Query(commentID, commentID, startNum, perPage)
	} else {
		rows, err = preGetReplies.Query(commentID, startNum, perPage)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return iterRowToList(rows)
}

func iterRowToList(rows *sql.Rows) (*List, error) {
	l := NewList()
	var (
		i   int64
		err error
	)

	for rows.Next() {
		c := newComment()
		err = rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.ParentID, &c.Created, &c.Updated, &c.Content)
		if err != nil {
			return nil, err
		}

		if err = c.CountReplies(); err != nil {
			return nil, err
		}

		l.Comments[c.ID] = c
		if c.ParentID != 0 {
			if l.Comments[c.ParentID] != nil {
				if len(l.Comments[c.ParentID].Children) == 0 {
					l.Comments[c.ParentID].Children[0] = c.ID
				} else {
					l.Comments[c.ParentID].Children[int64(len(l.Comments[c.ParentID].Children)-1)] = c.ID
				}
			}
		} else {
			l.Order[i] = c.ID
			i++
		}
	}

	return l, nil
}
