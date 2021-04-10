package avatar

import (
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/random"
)

func TransformAndSave(userId int64, tmpImagePath string) (newUrl string, err error) {
	img, err := imagedit.NewFromPath(tmpImagePath)
	if err != nil {
		return
	}
	defer img.Close()

	img.Strip()
	// Allow Animations
	img.TransformAnimation = true

	imgPath := locationPath(userId, random.StringBase58(16))
	fileImagePath := filePath(imgPath)

	w, h := img.AR(imagedit.AR_1x1)
	img.Crop(w, h, imagedit.Center)
	img.Resize(MaxHightWidth, MaxHightWidth)
	img.Optimize()

	img.WriteToPath(fileImagePath)

	err = InsertAvatar(userId, imgPath)
	if err == nil {
		newUrl = g.Config.Path.Public.Media + imgPath
	}
	return
}
