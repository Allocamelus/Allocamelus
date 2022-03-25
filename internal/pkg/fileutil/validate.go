package fileutil

import (
	"errors"
	"mime/multipart"
	"net/http"
	"sort"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

type VaidateConfig struct {
	MaxSize      int64
	ContentTypes []string
}

var (
	ErrContentType        = errors.New("invalid-content-type")
	ErrSomethingWentWrong = errors.New(apierr.SomethingWentWrong.String())
	ErrFileSize           = errors.New("invalid-file-size")
)

var (
	ImageContentTypes = []string{"image/png", "image/jpeg", "image/gif", "image/webp"}
)

func NewVaidateConfig(maxSize int64, contentTypes ...string) *VaidateConfig {
	vc := new(VaidateConfig)
	vc.MaxSize = maxSize
	sort.Strings(contentTypes)
	vc.ContentTypes = contentTypes
	return vc
}

func (vc *VaidateConfig) MpFileHeader(fileHead *multipart.FileHeader) error {
	if err := vc.Size(fileHead.Size); err != nil {
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

	if err := vc.ContentType(http.DetectContentType(buff)); err != nil {
		return err
	}

	return nil
}

func (vc *VaidateConfig) ContentType(contentType string) error {
	i := sort.SearchStrings(vc.ContentTypes, contentType)
	if len(vc.ContentTypes) == i {
		return ErrContentType
	}
	return nil
}

func (vc *VaidateConfig) Size(size int64) error {
	if size > vc.MaxSize {
		return ErrFileSize
	}
	return nil
}
