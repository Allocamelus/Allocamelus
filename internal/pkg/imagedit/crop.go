package imagedit

import "github.com/discord/lilliput"

/* not supported by lilliput ImageOptions
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
*/
func (img *Image) Crop(newWidth, newHeight int) {
	img.options.Width = newWidth
	img.options.Height = newHeight
	img.options.ResizeMethod = lilliput.ImageOpsFit
}

func (img *Image) CropAR(ar AspectRatio) {
	newWidth, newHeight := img.AR(ar)
	img.Crop(newWidth, newHeight)
}

/*
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
*/
