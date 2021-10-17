package imagedit

import (
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

func strToFmt(f string) fileutil.Format {
	switch f {
	case "GIF", "gif":
		return fileutil.GIF
	case "JPEG", "jpeg", "JPG", "jpg":
		return fileutil.JPG
	case "PNG", "png":
		return fileutil.PNG
	case "WEBP", "webp":
		return fileutil.WEBP
	}
	return fileutil.NONE
}

func (img *Image) GetFormat() fileutil.Format {
	return fileutil.ExtensionToFormat(img.options.FileType)
}
