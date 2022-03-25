package media

import (
	"mime/multipart"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

const (
	MaxImageSize int64 = 10 * 1024 * 1024 // 10Mb
)

var fileValidator = fileutil.NewVaidateConfig(MaxImageSize, fileutil.ImageContentTypes...)

// ValidateMpFileHeader multipart.FileHeader
func ValidateMpFileHeader(fileHead *multipart.FileHeader) error {
	return fileValidator.MpFileHeader(fileHead)
}
