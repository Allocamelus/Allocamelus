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

	err = img.Strip()
	if err != nil {
		return
	}
	// Allow Animations
	img.TransformAnimation = true

	err = img.CropAR(imagedit.AR_1x1, imagedit.Center)
	if err != nil {
		return
	}

	if err = img.Resize(MaxHightWidth, MaxHightWidth); err != nil {
		return
	}

	if err = img.Optimize(); err != nil {
		return
	}

	if err = InsertAvatar(userId, imgType, b58hash); err != nil {
		return
	}
	imgPath := selectorPath(b58hash, imgType, true)
	fileImagePath := fileutil.FilePath(imgPath)

	logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(b58hash, imgType, false))))

	err = img.WriteToPath(fileImagePath)
	if err != nil {
		return
	}

	logger.Error(deactivateOld(userId))
	newUrl = fileutil.PublicPath(imgPath)
	return
}
