package fileutil

import (
	"errors"
	"os"
)

func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
