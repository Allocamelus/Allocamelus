package dirutil

import (
	"os"
	"path/filepath"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
)

// RandomTmpDir create unique tmp path
func RandomTmpDir() (path string) {
	for {
		// Random tmp dir
		path = filepath.Join(g.Config.Path.TmpDir, random.StringBase64(8))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			logger.Error(os.MkdirAll(path, os.ModeSticky|os.ModePerm))
			break
		}
	}
	return path
}
