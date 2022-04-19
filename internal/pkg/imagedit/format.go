package imagedit

import (
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

func (img *Image) GetFormat() fileutil.Format {
	return fileutil.ExtensionToFormat(img.img.Format().FileExt())
}
