package avatar

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"gopkg.in/gographics/imagick.v3/imagick"
)

var MaxHightWidthStr = strconv.Itoa(int(MaxHightWidth))

func TransformAndSave(userId int64, tmpImagePath string) (newUrl string, err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	start := time.Now()
	slice := start

	err = mw.ReadImage(tmpImagePath)
	if err != nil {
		return
	}
	fmt.Printf("Read: %v\n", time.Since(slice))

	imgPath := "users/avatars/" + strconv.Itoa(int(userId)) + "/" + random.StringBase58(16)
	fileImagePath := filepath.Join(g.Config.Path.Media, imgPath)

	// check for animation
	if mw.GetNumberImages() > 1 {
		slice = time.Now()
		ret, err := imagick.ConvertImageCommand([]string{
			"convert", tmpImagePath, "-coalesce", "-resize", MaxHightWidthStr + "x" + MaxHightWidthStr, "-layers", "Optimize", fileImagePath,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("convert: %v\n", time.Since(slice))

		fmt.Printf("Metadata:\n%s\n", ret.Meta)

		slice = time.Now()
		delay := mw.GetImageDelay()
		aw := mw.CoalesceImages()
		mw.Destroy()
		defer aw.Destroy()

		mw = imagick.NewMagickWand()
		mw.SetImageDelay(delay)
		// deferred ln:16
		for i := 0; i < int(aw.GetNumberImages()); i++ {
			aw.SetIteratorIndex(i)
			img := aw.GetImage()

			width, height, newHW := imgSquare(img)
			cropAndResize(width, height, newHW, userId, img)
			mw.AddImage(img)
			img.Destroy()
		}
		fmt.Printf("convert: %v\n", time.Since(slice))

		slice = time.Now()
		//mw = mw.OptimizeImageLayers()
		fmt.Printf("OptimizeImageLayers: %v\n", time.Since(slice))
		slice = time.Now()
		mw.WriteImages(fileImagePath, true)
		fmt.Printf("Write: %v\n", time.Since(slice))
	} else {
		mw.StripImage()
		width, height, newHW := imgSquare(mw)
		err = cropAndResize(width, height, newHW, userId, mw)
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
		// Log to catch abuse
		logger.Error(errors.New("imagemagick: Error from User " + strconv.Itoa(int(userId)) + " | " + err.Error()))
		return err
	}

	err = mw.ResizeImage(MaxHightWidth, MaxHightWidth, imagick.FILTER_LANCZOS2)
	if err != nil {
		// Log to catch abuse
		logger.Error(errors.New("imagemagick: Error from User " + strconv.Itoa(int(userId)) + " | " + err.Error()))
		return err
	}
	return nil
}
