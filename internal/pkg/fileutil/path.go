package fileutil

import (
	"path/filepath"

	"github.com/allocamelus/allocamelus/internal/g"
)

func FilePath(relativePath string) string {
	return filepath.Join(g.Config.Path.Media, relativePath)
}

func PublicPath(relativePath string) string {
	return filepath.Join(g.Config.Path.Public.Media, relativePath)
}
