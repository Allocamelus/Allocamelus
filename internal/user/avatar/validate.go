package avatar

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

const (
	maxAvatarSize int64 = 1024 * 1024 * 5 // 5Mb
)

var (
	ErrContentType        = errors.New("invalid-content-type")
	ErrSomethingWentWrong = errors.New(apierr.SomethingWentWrong.String())
	ErrFileSize           = errors.New("invalid-file-size")
)

// ValidateMpFileHeader multipart.FileHeader
func ValidateMpFileHeader(fileHead *multipart.FileHeader) error {
	if err := ValidateSize(fileHead.Size); err != nil {
		return err
	}

	file, err := fileHead.Open()
	if logger.Error(err) {
		return ErrSomethingWentWrong
	}
	defer file.Close()

	buff := make([]byte, 512)
	if _, err = file.Read(buff); logger.Error(err) {
		return ErrSomethingWentWrong
	}

	if err := ValidateContentType(http.DetectContentType(buff)); err != nil {
		return err
	}

	return nil
}

func ValidateContentType(contentType string) error {
	switch contentType {
	case "image/png", "image/jpeg", "image/gif", "image/webp":
		return nil
	default:
		return ErrContentType
	}
}

func ValidateSize(size int64) error {
	if size > maxAvatarSize {
		return ErrFileSize
	}
	return nil
}
