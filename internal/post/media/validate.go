package media

import (
	"mime/multipart"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

const (
	MaxImageSize int64 = 10 * 1024 * 1024 // 10Mb
)

var (
	ContentTypes  = []string{"image/png", "image/jpeg", "image/gif", "image/webp"}
	fileValidator = fileutil.NewVaidateConfig(MaxImageSize, ContentTypes...)
)

// ValidateMpFileHeader multipart.FileHeader
func ValidateMpFileHeader(fileHead *multipart.FileHeader) error {
	return fileValidator.MpFileHeader(fileHead)
}
