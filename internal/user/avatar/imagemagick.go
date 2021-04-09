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

const (
	WEBP = "WEBP"
	PNG  = "PNG"
	JPG  = "JPEG"
	GIF  = "GIF"
)

func TransformAndSave(userId int64, tmpImagePath string) (newUrl string, err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err = mw.ReadImage(tmpImagePath)
	if err != nil {
		return
	}

	imgPath := "users/avatars/" + strconv.Itoa(int(userId)) + "/" + random.StringBase58(16)
	fileImagePath := filepath.Join(g.Config.Path.Media, imgPath)

	mw.StripImage()

	// check for animation
	if mw.GetNumberImages() > 1 {
		mw, err = transformAnimation(userId, mw)
		if err != nil {
			return
		}
		mw.WriteImages(fileImagePath, true)
	} else {
		err = transformImage(userId, mw)
		if err != nil {
			return
		}
		mw.WriteImage(fileImagePath)
	}

	err = InsertAvatar(userId, imgPath)
	if err == nil {
		newUrl = g.Config.Path.Public.Media + imgPath
	}
	return
}

func imgSquare(mw *imagick.MagickWand) (width, height, newHW uint) {
	width = mw.GetImageWidth()
	height = mw.GetImageHeight()
	if width > height {
		newHW = height
	} else {
		newHW = width
	}
	return
}

func cropAndResize(width, height, newHW uint, userId int64, mw *imagick.MagickWand) error {
	err := mw.CropImage(newHW, newHW, int((width-newHW)/2), int((height-newHW)/2))
	if err != nil {
		if userId != 0 {
			// Log to catch abuse
			logger.Error(errors.New("imagemagick: Error from User " + strconv.Itoa(int(userId)) + " | " + err.Error()))
		}
		return err
	}

	// resize if bigger
	if newHW > MaxHightWidth {
		err = mw.ResizeImage(MaxHightWidth, MaxHightWidth, imagick.FILTER_LANCZOS2)
		if err != nil {
			if userId != 0 {
				// Log to catch abuse
				logger.Error(errors.New("imagemagick: Error from User " + strconv.Itoa(int(userId)) + " | " + err.Error()))
			}
			return err
		}
	}

	return nil
}

func transformAnimation(userId int64, mw *imagick.MagickWand) (*imagick.MagickWand, error) {
	delay := mw.GetImageDelay()
	aw := mw.CoalesceImages()
	mw.Destroy()
	defer aw.Destroy()

	nmw := imagick.NewMagickWand()
	nmw.SetImageDelay(delay)
	// deferred ln:16
	for i := 0; i < int(aw.GetNumberImages()); i++ {
		aw.SetIteratorIndex(i)
		img := aw.GetImage()
		err := transformImage(userId, img)
		if err != nil {
			return nmw, err
		}
		nmw.AddImage(img)
		img.Destroy()
	}
	if nmw.GetImageFormat() == GIF {
		nmw.OptimizeImageLayers()
	}
	return nmw, nil
}

func transformImage(userId int64, mw *imagick.MagickWand) error {
	width, height, newHW := imgSquare(mw)
	return cropAndResize(width, height, newHW, userId, mw)
}
