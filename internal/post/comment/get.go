package comment

import (
	"context"
	_ "embed"
	"errors"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jackc/pgx/v5"
)

// Get
func Get(commentId int64) (*Comment, error) {
	dbc, err := g.Data.Queries.GetPostComment(context.Background(), commentId)
	if err != nil {
		return nil, err
	}
	c := DBCommentToComment(&dbc.Postcomment)
	c.Depth = dbc.Depth

	// Get reply count if any
	if err := c.CountReplies(); err != nil {
		return nil, err
	}

	return c, err
}

// GetForUser
func GetForUser(commentId int64, u *session.Session) (*Comment, error) {
	c, err := Get(commentId)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return c, err
		}
		return nil, ErrNoComment
	}

	if err = CanView(commentId, u, c); err != nil {
		return nil, err
	}

	return c, nil
}

func GetUserId(commentID int64) (int64, error) {
	return g.Data.Queries.GetPostCommentUserID(context.Background(), commentID)
}

func GetPostID(commentID int64) (int64, error) {
	// Get comment from store
	postID, err := g.Data.Queries.GetPostCommentPostID(context.Background(), commentID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return 0, err
		}
		return 0, ErrNoComment
	}
	return postID, nil
}

func GetPostUserID(commentID int64) (*Comment, error) {
	// Get comment from store
	c := newComment()
	c.PostID = commentID
	dbc, err := g.Data.Queries.GetPostCommentPostUserID(context.Background(), commentID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, ErrNoComment
	}
	c.PostID = dbc.Postid
	c.UserID = dbc.Userid
	return c, nil
}

// GetTotal comments for post
func GetPostTotal(postID int64, depth ...int64) (total int64, err error) {
	if len(depth) > 0 {
		total, err = g.Data.Queries.CountPostCommentsTotalDepth(context.Background(), db.CountPostCommentsTotalDepthParams{
			Postid: postID,
			Depth:  depth[0],
		})
		return
	}

	total, err = g.Data.Queries.CountPostCommentsTotal(context.Background(), postID)
	return
}

// GetPostTopLevel Get total top level (not a reply) comments for a post
func GetPostTopLevel(postID int64) (total int64, err error) {
	return g.Data.Queries.CountPostCommentsTopLevel(context.Background(), postID)
}

// GetPostComments
//
// topPerPage maximum num of top level comments per page
func GetPostComments(startNum, topPerPage, maxPerPage, postID int64, depth int64) (*List, error) {
	//rows, err := preGetPostComments.Query(postID, startNum, topPerPage, depth, maxPerPage)
	rows, err := g.Data.Queries.GetPostComments(context.Background(), db.GetPostCommentsParams{Postid: postID, Offset: int32(startNum), Limit: int32(topPerPage), Depth: depth, Limit_2: int32(maxPerPage)})
	if err != nil {
		return nil, err
	}

	return iterRowToList(rows, false)
}

// GetRepliesTotal
func GetRepliesTotal(commentID int64, depth ...int64) (total int64, err error) {
	if len(depth) > 0 {
		total, err = g.Data.Queries.CountPostCommentRepliesTotalDepth(context.Background(), db.CountPostCommentRepliesTotalDepthParams{Parent: commentID, Depth: depth[0]})
		return
	}

	total, err = g.Data.Queries.CountPostCommentReplies(context.Background(), commentID)
	return
}

func GetReplies(startNum, perPage, commentID, depth int64) (*List, error) {
	rows, err := g.Data.Queries.GetPostCommentReplies(context.Background(), db.GetPostCommentRepliesParams{Parent: commentID, Depth: depth, Offset: int32(startNum), Limit: int32(perPage)})

	if err != nil {
		return nil, err
	}
	return iterRowToList(rows, true)
}

func iterRowToList[R []db.GetPostCommentRepliesRow | []db.GetPostCommentsRow](rows R, noParents bool) (*List, error) { // skipcq: RVV-A0005
	l := NewList()
	var i int64

	locationMap := CommentList{}

	processComment := func(c *Comment) error {
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
		return nil
	}

	switch v := any(rows).(type) {
	case []db.GetPostCommentRepliesRow:
		for _, r := range v {
			c := DBCommentToComment(&r.Postcomment)
			c.Depth = r.Depth

			if err := processComment(c); err != nil {
				return nil, err
			}
		}
	case []db.GetPostCommentsRow:
		for _, r := range v {
			c := DBCommentToComment(&r.Postcomment)
			c.Depth = r.Depth

			if err := processComment(c); err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New("unknown row type in iterRowToList")
	}

	return l, nil
}
