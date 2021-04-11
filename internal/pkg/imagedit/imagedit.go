package imagedit

import (
	"errors"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type Image struct {
	// MagickWand
	MW                 *imagick.MagickWand
	Animation          bool
	TransformAnimation bool
	// OptimizeImageLayers panics if images are not all the same size
	resized bool
}

const (
	WEBP = "WEBP"
	PNG  = "PNG"
	JPG  = "JPEG"
	GIF  = "GIF"
)

var ErrNilWand = errors.New("imagedit: Error Nil MagickWand")

func NewFromPath(imagePath string) (*Image, error) {
	mw := imagick.NewMagickWand()

	err := mw.ReadImage(imagePath)
	if err != nil {
		return nil, err
	}
	return NewFromMW(mw)
}

func NewFromMW(mw *imagick.MagickWand) (*Image, error) {
	if mw == nil {
		return nil, ErrNilWand
	}

	img := new(Image)
	img.MW = mw
	img.checkAnimation()

	return img, nil
}

func (img *Image) WriteToPath(imagePath string) error {
	if err := img.Check(); err != nil {
		return err
	}
	if img.Animation {
		return img.MW.WriteImages(imagePath, true)
	}
	return img.MW.WriteImage(imagePath)
}

func (img *Image) NewMW(mw ...*imagick.MagickWand) {
	img.Close()
	if len(mw) > 0 {
		if mw[0] != nil {
			img.MW = mw[0]
			return
		}
	}
	img.MW = imagick.NewMagickWand()
}

func (img *Image) Check() error {
	if img.MW == nil {
		return ErrNilWand
	}
	return nil
}

func (img *Image) Close() {
	if err := img.Check(); err != nil {
		return
	}
	img.MW.Destroy()
}

// WH returns image width & height
//  return width uint, height uint
func (img *Image) WH() (width, height uint) {
	width = img.MW.GetImageWidth()
	height = img.MW.GetImageHeight()
	return
}
