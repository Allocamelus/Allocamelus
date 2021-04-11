package avatar

import (
	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
)

func TransformAndSave(userId int64, tmpImagePath string) (newUrl string, err error) {
	img, err := imagedit.NewFromPath(tmpImagePath)
	if err != nil {
		return
	}
	defer img.Close()

	selector := random.StringBase58(16)

	avatarId, err := InsertAvatar(userId, selector)
	if err != nil {
		return
	}

	fileImagePath := filePath(avatarId, selector)

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
	err = img.Resize(MaxHightWidth, MaxHightWidth)
	if err != nil {
		return
	}
	err = img.Optimize()
	if err != nil {
		return
	}

	logger.Error(dirutil.MakeDir(filePath(avatarId, "")))

	err = img.WriteToPath(fileImagePath)
	if err != nil {
		return
	}

	logger.Error(deactivateOld(userId, avatarId))
	newUrl = publicPath(avatarId, selector)
	return
}
