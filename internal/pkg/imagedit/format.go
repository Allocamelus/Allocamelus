package imagedit

import "github.com/allocamelus/allocamelus/internal/pkg/fileutil"

func mwToFmt(f string) fileutil.Format {
	switch f {
	case "GIF":
		return fileutil.GIF
	case "JPEG":
		return fileutil.JPG
	case "PNG":
		return fileutil.PNG
	case "WEBP":
		return fileutil.WEBP
	}
	return fileutil.NONE
}

func (img *Image) GetFormat() fileutil.Format {
	return mwToFmt(img.MW.GetImageFormat())
}
