package comment

import (
	"database/sql"
	_ "embed"
	"errors"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

var (
	//go:embed sql/get.sql
	qGet   string
	preGet *sql.Stmt
)

// Get
func Get(commentId int64) (*Comment, error) {
	if preGet == nil {
		preGet = g.Data.Prepare(qGet)
	}
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
	//go:embed sql/getUserID.sql
	qGetUserID   string
	preGetUserID *sql.Stmt
)

func GetUserId(commentID int64) (int64, error) {
	if preGetUserID == nil {
		preGetUserID = g.Data.Prepare(qGetUserID)
	}
	var userId int64
	err := preGetUserID.QueryRow(commentID).Scan(&userId)
	return userId, err
}

var (
	//go:embed sql/getForCanView.sql
	qGetForCanView string
	preGetCanView  *sql.Stmt
)

func getForCanView(commentID int64) (*Comment, error) {
	if preGetCanView == nil {
		preGetCanView = g.Data.Prepare(qGetForCanView)
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

var (
	//go:embed sql/getPostTotal.sql
	qGetPostTotal string
	//go:embed sql/getPostTotalDepth.sql
	qGetPostTotalDepth   string
	preGetPostTotal      *sql.Stmt
	preGetPostTotalDepth *sql.Stmt
)

// GetTotal comments for post
func GetPostTotal(postID int64, depth ...int64) (total int64, err error) {
	if preGetPostTotal == nil {
		preGetPostTotal = g.Data.Prepare(qGetPostTotal)
	}
	if preGetPostTotalDepth == nil {
		preGetPostTotalDepth = g.Data.Prepare(qGetPostTotalDepth)
	}

	if len(depth) > 0 {
		err = preGetPostTotalDepth.QueryRow(postID, depth[0]).Scan(&total)
		return
	}

	err = preGetPostTotal.QueryRow(postID).Scan(&total)
	return
}

var (
	//go:embed sql/getPostTopLevel.sql
	qGetPostTopLevel   string
	preGetPostTopLevel *sql.Stmt
)

// GetPostTopLevel Get total top level (not a reply) comments for a post
func GetPostTopLevel(postID int64) (total int64, err error) {
	if preGetPostTopLevel == nil {
		preGetPostTopLevel = g.Data.Prepare(qGetPostTopLevel)
	}
	err = preGetPostTopLevel.QueryRow(postID).Scan(&total)
	return
}

// Get Post comment queries
var (
	//go:embed sql/getPostComments.sql
	qGetPostComments   string
	preGetPostComments *sql.Stmt
)

func GetPostComments(startNum, perPage, postID int64, depth int64) (*List, error) {
	if preGetPostComments == nil {
		preGetPostComments = g.Data.Prepare(qGetPostComments)
	}

	rows, err := preGetPostComments.Query(postID, startNum, perPage, depth, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return iterRowToList(rows, false)
}

var (
	//go:embed sql/getRepliesTotal.sql
	qGetRepliesTotal string
	//go:embed sql/getRepliesTotalDepth.sql
	qGetRepliesTotalDepth   string
	preGetRepliesTotal      *sql.Stmt
	preGetRepliesTotalDepth *sql.Stmt
)

// GetRepliesTotal
func GetRepliesTotal(commentID int64, depth ...int64) (total int64, err error) {
	// PCC.parent != PCC.child prevents query from counting parent's self
	if preGetRepliesTotal == nil {
		preGetRepliesTotal = g.Data.Prepare(qGetRepliesTotal)
	}
	if preGetRepliesTotalDepth == nil {
		preGetRepliesTotalDepth = g.Data.Prepare(qGetRepliesTotalDepth)
	}

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
	//go:embed sql/getReplies.sql
	qGetReplies   string
	preGetReplies *sql.Stmt
)

func GetReplies(startNum, perPage, commentID int64, depth int64) (*List, error) {
	if preGetReplies == nil {
		preGetReplies = g.Data.Prepare(qGetReplies)
	}

	rows, err := preGetReplies.Query(commentID, depth, startNum, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return iterRowToList(rows, true)
}

func iterRowToList(rows *sql.Rows, noParents bool) (*List, error) {
	l := NewList()
	var (
		i   int64
		err error
	)

	locationMap := map[int64]*Comment{}

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
