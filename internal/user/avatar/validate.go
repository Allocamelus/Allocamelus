package avatar

import (
	"mime/multipart"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
)

const (
	MaxAvatarSize int64 = 1024 * 1024 * 5 // 5Mb
)

var fileValidator = fileutil.NewVaidateConfig(MaxAvatarSize, fileutil.ImageContentTypes...)

// ValidateMpFileHeader multipart.FileHeader
func ValidateMpFileHeader(fileHead *multipart.FileHeader) error {
	return fileValidator.MpFileHeader(fileHead)
}
