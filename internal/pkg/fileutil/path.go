package fileutil

import (
	"path/filepath"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
)

func FilePath(relativePath string) string {
	return filepath.Join(g.Config.Path.MediaDir, relativePath)
}

func PublicPath(relativePath string) string {
	return filepath.Join(g.Config.Path.Public.Media, relativePath)
}

func RelativePath(prePath, b58hash string, includeFile bool) string {
	var path strings.Builder
	path.WriteString(prePath + "/" + b58hash[:3] + "/" + b58hash[3:6])
	if includeFile {
		path.WriteString("/" + b58hash)
	}
	return path.String()
}
