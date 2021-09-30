package comment

import (
	"database/sql"
	_ "embed"
	"errors"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

var (
	//go:embed sql/get/get.sql
	qGet   string
	preGet *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGet, qGet)
}

// Get
func Get(commentId int64) (*Comment, error) {
	c := newComment()
	c.ID = commentId
	err := preGet.QueryRow(commentId).Scan(&c.PostID, &c.UserID, &c.ParentID, &c.Created, &c.Updated, &c.Content, &c.Depth)
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

var (
	//go:embed sql/get/userID.sql
	qGetUserID   string
	preGetUserID *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetUserID, qGetUserID)
}

func GetUserId(commentID int64) (int64, error) {
	var userId int64
	err := preGetUserID.QueryRow(commentID).Scan(&userId)
	return userId, err
}

var (
	//go:embed sql/get/postID.sql
	qGetPostID   string
	preGetPostID *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetPostID, qGetPostID)
}

func GetPostID(commentID int64) (int64, error) {
	// Get comment from store
	var postID int64
	err := preGetPostID.QueryRow(commentID).Scan(&postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
		return 0, ErrNoComment
	}
	return postID, nil
}

var (
	//go:embed sql/get/postUserID.sql
	qGetPostUserID   string
	preGetPostUserID *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetPostUserID, qGetPostUserID)
}

func GetPostUserID(commentID int64) (*Comment, error) {
	// Get comment from store
	c := newComment()
	c.PostID = commentID
	err := preGetPostUserID.QueryRow(commentID).Scan(&c.PostID, &c.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, ErrNoComment
	}
	return c, nil
}

var (
	//go:embed sql/get/postTotal.sql
	qGetPostTotal   string
	preGetPostTotal *sql.Stmt
	//go:embed sql/get/postTotalDepth.sql
	qGetPostTotalDepth   string
	preGetPostTotalDepth *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetPostTotal, qGetPostTotal)
	data.PrepareQueuer.Add(&preGetPostTotalDepth, qGetPostTotalDepth)
}

// GetTotal comments for post
func GetPostTotal(postID int64, depth ...int64) (total int64, err error) {
	if len(depth) > 0 {
		err = preGetPostTotalDepth.QueryRow(postID, depth[0]).Scan(&total)
		return
	}

	err = preGetPostTotal.QueryRow(postID).Scan(&total)
	return
}

var (
	//go:embed sql/get/postTopLevel.sql
	qGetPostTopLevel   string
	preGetPostTopLevel *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetPostTopLevel, qGetPostTopLevel)
}

// GetPostTopLevel Get total top level (not a reply) comments for a post
func GetPostTopLevel(postID int64) (total int64, err error) {
	err = preGetPostTopLevel.QueryRow(postID).Scan(&total)
	return
}

// Get Post comment queries
var (
	//go:embed sql/get/postComments.sql
	qGetPostComments   string
	preGetPostComments *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetPostComments, qGetPostComments)
}

// GetPostComments
//
// topPerPage maximum num of top level comments per page
func GetPostComments(startNum, topPerPage, maxPerPage, postID int64, depth int64) (*List, error) {
	rows, err := preGetPostComments.Query(postID, startNum, topPerPage, depth, maxPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return iterRowToList(rows, false)
}

var (
	//go:embed sql/get/repliesTotal.sql
	qGetRepliesTotal   string
	preGetRepliesTotal *sql.Stmt
	//go:embed sql/get/repliesTotalDepth.sql
	qGetRepliesTotalDepth   string
	preGetRepliesTotalDepth *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetRepliesTotal, qGetRepliesTotal)
	data.PrepareQueuer.Add(&preGetRepliesTotalDepth, qGetRepliesTotalDepth)
}

// GetRepliesTotal
func GetRepliesTotal(commentID int64, depth ...int64) (total int64, err error) {
	var row *sql.Row

	if len(depth) > 0 {
		row = preGetRepliesTotalDepth.QueryRow(commentID, depth[0])
	} else {
		row = preGetRepliesTotal.QueryRow(commentID)
	}

	err = row.Scan(&total)
	return
}

var (
	//go:embed sql/get/replies.sql
	qGetReplies   string
	preGetReplies *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetReplies, qGetReplies)
}

func GetReplies(startNum, perPage, commentID int64, depth int64) (*List, error) {
	rows, err := preGetReplies.Query(commentID, depth, startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return iterRowToList(rows, true)
}

func iterRowToList(rows *sql.Rows, noParents bool) (*List, error) { // skipcq: RVV-A0005
	l := NewList()
	var (
		i   int64
		err error
	)

	locationMap := CommentList{}

	for rows.Next() {
		c := newComment()
		err = rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.ParentID, &c.Created, &c.Updated, &c.Content, &c.Depth)
		if err != nil {
			return nil, err
		}

		if err = c.CountReplies(); err != nil {
			return nil, err
		}

		// Map with pointers to all comments
		locationMap[c.ID] = c

		// If not a top level comment And (rows provides parents Or depth is > 1)
		if c.ParentID != 0 && (!noParents || c.Depth > 1) {
			// Check locationMap for comment
			if locationMap[c.ParentID] == nil {
				logger.Error(errors.New("comment/get: Error nil pointer to parent: " + strconv.Itoa(int(c.ParentID)) + " building iterRowToList"))
			} else {
				if len(locationMap[c.ParentID].Children) == 0 {
					locationMap[c.ParentID].Children[0] = c
				} else {
					locationMap[c.ParentID].Children[int64(len(locationMap[c.ParentID].Children))] = c
				}
			}
		} else {
			l.Comments[c.ID] = c
			l.Order[i] = c.ID
			i++
		}
	}

	return l, nil
}
