//go:generate msgp

package media

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

type MediaType int8

const (
	None MediaType = iota
	Image
)

type Media struct {
	MediaType MediaType `msg:"mediaType" json:"mediaType"`
	Url       string    `msg:"url" json:"url"`
}

var (
	preGet    *sql.Stmt
	preInsert *sql.Stmt
)

func Get(postID int64) ([]*Media, error) {
	if preGet == nil {
		preGet = g.Data.Prepare(`SELECT postMediaId, mediaType, selector FROM PostMedia WHERE postId = ? AND active = 1 ORDER BY postMediaId ASC LIMIT 4`)
	}
	rows, err := preGet.Query(postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mediaList := []*Media{}

	for rows.Next() {
		var (
			mediaId   int64
			mediaType int8
			selector  string
		)
		err := rows.Scan(&mediaId, &mediaType, &selector)
		if err != nil {
			return nil, err
		}
		media := new(Media)
		media.MediaType = MediaType(mediaType)
		media.Url = fileutil.PublicPath(selectorPath(mediaId, selector, media.MediaType))
		mediaList = append(mediaList, media)
	}

	return mediaList, nil
}

func Insert(postID int64, mediaType MediaType, selector string) (int64, error) {
	if preInsert == nil {
		preInsert = g.Data.Prepare(`INSERT INTO PostMedia (postId, created, active, mediaType, selector) VALUES (?, ?, 1, ?, ?)`)
	}
	r, err := preInsert.Exec(postID, time.Now().Unix(), mediaType, selector)
	if err != nil {
		return 0, err
	}
	mediaId, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return mediaId, nil
}

func selectorPath(avatarId int64, selector string, mediaType MediaType) string {
	var sb strings.Builder
	sb.WriteRune('/')
	switch mediaType {
	case Image:
		sb.WriteString("images")
	}
	sb.WriteRune('/')
	return "posts/media/" + strconv.Itoa(int(avatarId)) + sb.String() + selector
}
