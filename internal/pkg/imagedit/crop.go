package imagedit

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
)

func (img *Image) Crop(newWidth, newHeight uint, l Location) error {
	if err := img.Check(); err != nil {
		return err
	}

	if img.Animation {
		callback := func(callbackImg *Image) error {
			return callbackImg.Crop(newWidth, newHeight, l)
		}
		return img.IterateOver(callback)
	}

	width, height := img.WH()
	x, y := fromLocation(int(width), int(height), int(newWidth), int(newHeight), l)

	return img.MW.CropImage(newWidth, newHeight, x, y)
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
