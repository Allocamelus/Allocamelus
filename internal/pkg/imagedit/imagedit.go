package imagedit

import (
	"errors"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/h2non/bimg"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type Image struct {
	// MagickWand
	MW    *imagick.MagickWand
	Img   *bimg.Image
	UseMW bool
	// OptimizeImageLayers panics if images are not all the same size
	resized bool
}

var (
	ErrNilWand = errors.New("imagedit: Error Nil MagickWand")
	ErrNilImg  = errors.New("imagedit: Error Nil bimg.Image")
)

func NewFromPath(imagePath string) (*Image, error) {
	imgBlob, err := bimg.Read(imagePath)
	if err != nil {
		return nil, err
	}
	return NewFromBlob(imgBlob)
}

func NewFromBlob(blob []byte) (*Image, error) {
	img := new(Image)
	img.BlobToImg(blob, nil)
	// h2non/bimg doesn't have gif or animation save support
	imgFmt := img.GetFormat()
	if imgFmt == fileutil.GIF || imgFmt == fileutil.WEBP {
		mw := imagick.NewMagickWand()
		err := mw.ReadImageBlob(blob)
		if err != nil {
			return nil, err
		}

		if isAnimation(mw) {
			img.MW = mw
			img.UseMW = true
			img.Strip()
			img.Img = nil
		} else {
			mw.Destroy()
			if imgFmt == fileutil.GIF {
				if err = img.BlobToImg(img.Img.Convert(bimg.PNG)); err != nil {
					return nil, err
				}
			}
		}
	}
	return img, nil
}

func (img *Image) BlobToImg(blob []byte, err error) error {
	img.Img = bimg.NewImage(blob)
	return err
}

func NewFromMW(mw *imagick.MagickWand) (*Image, error) {
	if mw == nil {
		return nil, ErrNilWand
	}

	img := new(Image)
	img.MW = mw
	img.UseMW = true

	return img, nil
}

func (img *Image) WriteToPath(imagePath string) error {
	if err := img.Check(); err != nil {
		return err
	}
	if img.UseMW {
		img.Optimize()
		return img.MW.WriteImages(imagePath, true)
	}
	return bimg.Write(imagePath, img.Img.Image())
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
	if img.UseMW {
		if img.MW == nil {
			return ErrNilImg
		}
	} else if img.Img == nil {
		return ErrNilImg
	}
	return nil
}

func (img *Image) Close() {
	if img.UseMW {
		if err := img.Check(); err != nil {
			return
		}
		img.MW.Destroy()
	}
}

// WH returns image width & height
//  return width uint, height uint
func (img *Image) WH() (width, height int, err error) {
	if err = img.Check(); err != nil {
		return
	}

	if img.UseMW {
		width = int(img.MW.GetImageWidth())
		height = int(img.MW.GetImageHeight())
		return
	}

	size, err := img.Img.Size()
	width = size.Width
	height = size.Height
	return
}
