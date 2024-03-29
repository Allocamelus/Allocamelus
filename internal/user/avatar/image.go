package avatar

import (
	"mime/multipart"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

func TransformAndSave(userId int64, imageMPH *multipart.FileHeader) (newUrl string, err error) {
	img, b58hash, err := imagedit.MPHtoImg(imageMPH)
	if err != nil {
		return
	}
	defer img.Close()

	imgType := img.GetFormat()
	if !imgType.IsImage() {
		err = fileutil.ErrContentType
		return
	}

	imgPath := selectorPath(b58hash, true)
	fileImagePath := fileutil.FilePath(imgPath)

	// Check for image for deduplication
	if !fileutil.Exist(fileImagePath) {
		err = img.Thumbnail(MaxHightWidth, MaxHightWidth)
		if err != nil {
			return
		}

		logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(b58hash, false))))

		err = img.WriteToPath(fileImagePath)
		if err != nil {
			return
		}
	}

	if err = InsertAvatar(userId, imgType, b58hash); err != nil {
		return
	}

	logger.Error(deactivateOld(userId))
	newUrl = fileutil.PublicPath(imgPath)
	return
}
