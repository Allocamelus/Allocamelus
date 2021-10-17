package media

import (
	"mime/multipart"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

const MaxHightWidth int = 8192

func TransformAndSave(postID int64, imageMPH *multipart.FileHeader, alt string) error {
	img, b58hash, err := imagedit.MPHtoImg(imageMPH)
	if err != nil {
		return err
	}
	defer img.Close()

	imgType := img.GetFormat()
	if !imgType.IsImage() {
		return fileutil.ErrContentType
	}

	fileImagePath := fileutil.FilePath(selectorPath(b58hash, true))

	// Check for image for deduplication
	if !fileutil.Exist(fileImagePath) {
		width, height := img.WH()
		if err != nil {
			return err
		}
		var newWidth, newHeight int
		if width > MaxHightWidth || height > MaxHightWidth {
			newWidth, newHeight = img.ARMaxSize(imagedit.AR_Image, MaxHightWidth)
			if err != nil {
				return err
			}
		} else {
			newWidth = width
			newHeight = height
		}
		img.Resize(newWidth, newHeight)

		logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(b58hash, false))))

		err = img.WriteToPath(fileImagePath)
		if err != nil {
			return err
		}
	} else {
		logger.Error(err)
	}

	width, height := img.WH()

	err = Insert(postID, Media{FileType: imgType, Meta: Meta{Alt: alt, Width: int64(width), Height: int64(height)}}, b58hash)
	if err != nil {
		return err
	}

	return err
}
