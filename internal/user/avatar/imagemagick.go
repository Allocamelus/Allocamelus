package avatar

import (
	"errors"
	"path/filepath"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func TransformAndSave(userId int64, tmpImagePath string) (newUrl string, err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err = mw.ReadImage(tmpImagePath)
	if err != nil {
		return
	}
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()
	var newHW uint
	if width > height {
		newHW = height
	} else {
		newHW = width
	}

	err = mw.CropImage(newHW, newHW, int((width-newHW)/2), int((height-newHW)/2))
	if err != nil {
		// Log to catch abuse
		logger.Error(errors.New("imagemagick: Error from User " + strconv.Itoa(int(userId)) + " | " + err.Error()))
		return
	}
	err = mw.ResizeImage(MaxHightWidth, MaxHightWidth, imagick.FILTER_LANCZOS2)
	if err != nil {
		// Log to catch abuse
		logger.Error(errors.New("imagemagick: Error from User " + strconv.Itoa(int(userId)) + " | " + err.Error()))
		return
	}

	imgPath := "users/avatars/" + strconv.Itoa(int(userId)) + "/" + random.StringBase58(16)
	mw.WriteImage(filepath.Join(g.Config.Path.Media, imgPath))
	err = InsertAvatar(userId, imgPath)
	if err == nil {
		newUrl = g.Config.Path.Public.Media + imgPath
	}
	return
}
