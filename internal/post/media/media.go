//go:generate msgp

package media

import (
	"database/sql"
	_ "embed"
	"html"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	jsoniter "github.com/json-iterator/go"
)

type Media struct {
	FileType fileutil.Format `msg:"fileType" json:"fileType,omitempty"`
	Meta     Meta            `msg:"meta" json:"meta"`
	Url      string          `msg:"url" json:"url"`
}

type Meta struct {
	Alt    string `msg:"alt" json:"alt"`
	Width  int64  `msg:"width" json:"width"`
	Height int64  `msg:"height" json:"height"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	//go:embed sql/get.sql
	qGet   string
	preGet *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGet, qGet)
}

func Get(postID int64) ([]*Media, error) {
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
		if err := rows.Scan(&meta, &hash); err != nil {
			return nil, err
		}
		if err := json.UnmarshalFromString(meta, &media.Meta); err != nil {
			return nil, err
		}
		media.Meta.Alt = html.EscapeString(media.Meta.Alt)
		media.Url = fileutil.PublicPath(selectorPath(hash, true))
		mediaList = append(mediaList, media)
	}

	return mediaList, nil
}

var (
	//go:embed sql/insert.sql
	qInsert   string
	preInsert *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preInsert, qInsert)
}

func Insert(postID int64, media Media, hash string) error {
	meta, err := json.MarshalToString(media.Meta)
	if err != nil {
		return err
	}
	_, err = preInsert.Exec(postID, time.Now().Unix(), media.FileType, meta, hash)
	if err != nil {
		return err
	}
	return nil
}

func selectorPath(b58hash string, includeFile bool) string {
	return fileutil.RelativePath("posts/images", b58hash, includeFile)
}
