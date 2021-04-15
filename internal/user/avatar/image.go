package avatar

import (
	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
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

	selector := random.StringBase58(16)

	avatarId, err := InsertAvatar(userId, selector)
	if err != nil {
		return
	}

	fileImagePath := fileutil.FilePath(selectorPath(avatarId, selector))

	logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(avatarId, ""))))

	err = img.WriteToPath(fileImagePath)
	if err != nil {
		return
	}

	logger.Error(deactivateOld(userId, avatarId))
	newUrl = fileutil.PublicPath(selectorPath(avatarId, selector))
	return
}
