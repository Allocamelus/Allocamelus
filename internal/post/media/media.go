//go:generate msgp

package media

import (
	"context"
	_ "embed"
	"html"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
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

func New() *Media {
	media := new(Media)
	media.Meta = new(Meta)
	return media
}

func Get(postID int64) ([]*Media, error) {
	rows, err := g.Data.Queries.GetPostMedia(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	mediaList := []*Media{}

	for _, r := range rows {
		media := New()
		var hash string
		media.Meta.Alt = html.EscapeString(r.Alt)
		media.Meta.Width = int64(r.Width)
		media.Meta.Height = int64(r.Height)
		hash = r.Hash

		media.Url = fileutil.PublicPath(selectorPath(hash, true))
		mediaList = append(mediaList, media)
	}

	return mediaList, nil
}

func Insert(postId int64, alt string, fileId int64) error {
	return g.Data.Queries.InsertPostMedia(context.Background(), db.InsertPostMediaParams{Postid: postId, Added: time.Now().Unix(), Alt: alt, Postmediafileid: fileId})
}

func InsertFile(media *Media, hash string, newHash string) (int64, error) {
	return g.Data.Queries.InsertPostMediaFile(context.Background(), db.InsertPostMediaFileParams{
		Created:  time.Now().Unix(),
		Filetype: int32(media.FileType),
		Width:    int32(media.Meta.Width),
		Height:   int32(media.Meta.Height),
		Hash:     hash,
		Newhash:  newHash,
	})
}

func selectorPath(b58hash string, includeFile bool) string {
	return fileutil.RelativePath(SubPath, b58hash, includeFile)
}
