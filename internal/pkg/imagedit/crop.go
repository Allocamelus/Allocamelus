package imagedit

import "github.com/h2non/bimg"

type Location int

const (
	Center Location = iota
	TopLeft
	Top
	TopRight
	Right
	BottomRight
	Bottom
	BottomLeft
	Left
	// bimg only mw defaults to center
	Smart
)

func (img *Image) Crop(newWidth, newHeight int, l Location) error {
	if err := img.Check(); err != nil {
		return err
	}

	if img.UseMW {
		callback := func(callbackImg *Image) error {
			return callbackImg.CropMW(newWidth, newHeight, l)
		}
		return img.IterateOver(callback)
	}

	return img.BlobToImg(img.Img.Crop(newWidth, newHeight, lToG(l)))
}

func (img *Image) CropMW(newWidth, newHeight int, l Location) error {
	if err := img.Check(); err != nil {
		return err
	}

	width, height, _ := img.WH()
	x, y := fromLocation(width, height, newWidth, newHeight, l)

	return img.MW.CropImage(uint(newWidth), uint(newHeight), x, y)
}

func (img *Image) CropAR(ar AspectRatio, l Location) error {
	if err := img.Check(); err != nil {
		return err
	}

	if img.UseMW {
		callback := func(callbackImg *Image) error {
			return callbackImg.CropArMw(ar, l)
		}
		return img.IterateOver(callback)
	}

	newWidth, newHeight, err := img.AR(ar)
	if err != nil {
		return err
	}
	return img.Crop(newWidth, newHeight, l)
}

func (img *Image) CropArMw(ar AspectRatio, l Location) error {
	if err := img.Check(); err != nil {
		return err
	}

	newWidth, newHeight, _ := img.AR(ar)
	return img.CropMW(newWidth, newHeight, l)
}

func fromLocation(w, h, newH, newW int, l Location) (x, y int) {
	switch l {
	case Center:
		return (w - newW) / 2, (h - newH) / 2
	case TopLeft:
		return 0, 0
	case Top:
		return (w - newW) / 2, 0
	case TopRight:
		return w - newW, 0
	case Right:
		return w - newW, (h - newH) / 2
	case BottomRight:
		return w - newW, h - newH
	case Bottom:
		return (w - newW) / 2, h - newH
	case BottomLeft:
		return 0, h - newH
	case Left:
		return 0, (h - newH) / 2
	}
	return (w - newW) / 2, (h - newH) / 2
}

func lToG(l Location) bimg.Gravity {
	switch l {
	case Top, TopLeft, TopRight:
		return bimg.GravityNorth
	case Right:
		return bimg.GravityEast
	case Bottom, BottomLeft, BottomRight:
		return bimg.GravitySouth
	case Left:
		return bimg.GravityWest
	case Smart:
		return bimg.GravitySmart
	}
	return bimg.GravityCentre
}
