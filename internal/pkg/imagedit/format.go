package imagedit

import (
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

func mwToFmt(f string) fileutil.Format {
	switch f {
	case "GIF", "gif":
		return fileutil.GIF
	case "JPEG", "jpeg":
		return fileutil.JPG
	case "PNG", "png":
		return fileutil.PNG
	case "WEBP", "webp":
		return fileutil.WEBP
	}
	return fileutil.NONE
}

func (img *Image) GetFormat() fileutil.Format {
	if img.UseMW {
		return mwToFmt(img.MW.GetImageFormat())
	}
	return mwToFmt(img.Img.Type())
}
