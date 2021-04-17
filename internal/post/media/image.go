package media

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

const MaxHightWidth uint = 7680

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

	fileImagePath := fileutil.FilePath(selectorPath(b58hash, imgType, true))

	_, err = os.Stat(fileImagePath)
	// Check for image for deduplication
	if errors.Is(err, os.ErrNotExist) {
		err = img.Strip()
		if err != nil {
			return err
		}
		// Allow Animations
		img.TransformAnimation = true
		width, height := img.WH()
		var newWidth, newHeight uint
		if width > MaxHightWidth || height > MaxHightWidth || img.Animation {
			newWidth, newHeight = img.ARMaxSize(imagedit.AR_Image, MaxHightWidth)
		} else {
			newWidth = width
			newHeight = height
		}
		// Resize to prevent non images
		if err = img.Resize(newWidth, newHeight); err != nil {
			return err
		}

		if err = img.Optimize(); err != nil {
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

	logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(b58hash, imgType, false))))

	err = img.WriteToPath(fileImagePath)
	if err != nil {
		return err
	}

	return err
}
