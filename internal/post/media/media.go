//go:generate msgp

package media

import (
	"database/sql"
	"html"
	"strings"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	jsoniter "github.com/json-iterator/go"
)

type MediaType int

const (
	None MediaType = iota
	IMG_PNG
	IMG_JPG
	IMG_WEBP
	IMG_GIF
)

type Media struct {
	MediaType MediaType `msg:"mediaType" json:"mediaType"`
	Meta      Meta      `msg:"meta" json:"meta"`
	Url       string    `msg:"url" json:"url"`
}

type Meta struct {
	Alt    string `msg:"alt" json:"alt"`
	Width  int64  `msg:"width" json:"width"`
	Height int64  `msg:"height" json:"height"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	preGet    *sql.Stmt
	preInsert *sql.Stmt
)

func Get(postID int64) ([]*Media, error) {
	if preGet == nil {
		preGet = g.Data.Prepare(`SELECT mediaType, meta, hash FROM PostMedia WHERE postId = ? AND active = 1 ORDER BY postMediaId ASC LIMIT 4`)
	}
	rows, err := preGet.Query(postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mediaList := []*Media{}

	for rows.Next() {
		media := new(Media)
		var (
			meta string
			hash string
		)
		if err := rows.Scan(&media.MediaType, &meta, &hash); err != nil {
			return nil, err
		}
		if err := json.UnmarshalFromString(meta, &media.Meta); err != nil {
			return nil, err
		}
		media.Meta.Alt = html.EscapeString(media.Meta.Alt)
		media.Url = fileutil.PublicPath(selectorPath(hash, media.MediaType, true))
		mediaList = append(mediaList, media)
	}

	return mediaList, nil
}

func Insert(postID int64, media Media, hash string) error {
	if preInsert == nil {
		preInsert = g.Data.Prepare(`INSERT INTO PostMedia (postId, created, active, mediaType, meta, hash) VALUES (?, ?, 1, ?, ?, ?)`)
	}
	meta, err := json.MarshalToString(media.Meta)
	if err != nil {
		return err
	}
	_, err = preInsert.Exec(postID, time.Now().Unix(), media.MediaType, meta, hash)
	if err != nil {
		return err
	}
	return nil
}

func ImageditFmtToType(f imagedit.Format) MediaType {
	switch f {
	case imagedit.GIF:
		return IMG_GIF
	case imagedit.JPG:
		return IMG_JPG
	case imagedit.PNG:
		return IMG_PNG
	case imagedit.WEBP:
		return IMG_WEBP
	}
	return None
}

func (mt MediaType) FileExt() string {
	switch mt {
	case IMG_GIF:
		return ".gif"
	case IMG_JPG:
		return ".jpg"
	case IMG_PNG:
		return ".png"
	case IMG_WEBP:
		return ".web"
	}
	return ""
}

func selectorPath(encodedHash string, mediaType MediaType, includeFile bool) string {
	var path strings.Builder
	path.WriteString("posts/" + encodedHash[:3] + "/" + encodedHash[3:6])
	if includeFile {
		path.WriteString("/" + encodedHash + mediaType.FileExt())
	}
	return path.String()
}
