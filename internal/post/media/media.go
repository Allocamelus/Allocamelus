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
	Meta     *Meta           `msg:"meta" json:"meta"`
	Url      string          `msg:"url" json:"url"`
}

type Meta struct {
	Alt    string `msg:"alt" json:"alt"`
	Width  int64  `msg:"width" json:"width"`
	Height int64  `msg:"height" json:"height"`
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	//go:embed sql/get/get.sql
	qGet   string
	preGet *sql.Stmt
	//go:embed sql/insert/media.sql
	qInsertMedia   string
	preInsertMedia *sql.Stmt
	//go:embed sql/insert/mediaFile.sql
	qInsertMediaFile   string
	preInsertMediaFile *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGet, qGet)
	data.PrepareQueuer.Add(&preInsertMediaFile, qInsertMediaFile)
	data.PrepareQueuer.Add(&preInsertMedia, qInsertMedia)
}

func New() *Media {
	media := new(Media)
	media.Meta = new(Meta)
	return media
}

func Get(postID int64) ([]*Media, error) {
	rows, err := preGet.Query(postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mediaList := []*Media{}

	for rows.Next() {
		media := New()
		var hash string
		if err := rows.Scan(&media.Meta.Alt, &media.Meta.Width, &media.Meta.Height, &hash); err != nil {
			return nil, err
		}
		media.Meta.Alt = html.EscapeString(media.Meta.Alt)
		media.Url = fileutil.PublicPath(selectorPath(hash, true))
		mediaList = append(mediaList, media)
	}

	return mediaList, nil
}

func Insert(postId int64, alt string, fileId int64) error {
	_, err := preInsertMedia.Exec(postId, time.Now().Unix(), alt, fileId)
	if err != nil {
		return err
	}
	return nil
}

func InsertFile(media *Media, hash string, newHash string) (int64, error) {
	n, err := preInsertMediaFile.Exec(time.Now().Unix(), media.FileType, media.Meta.Width, media.Meta.Height, hash, newHash)
	if err != nil {
		return 0, err
	}
	id, err := n.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func selectorPath(b58hash string, includeFile bool) string {
	return fileutil.RelativePath(SubPath, b58hash, includeFile)
}
