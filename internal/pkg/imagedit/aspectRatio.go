package imagedit

type AspectRatio int

const (
	// Use Image AspectRatio
	AR_Image AspectRatio = iota
	// 1:1 square
	AR_1x1
	// 5:4
	AR_5x4
	// 4:3
	AR_4x3
	// 3:2
	AR_3x2
	// 16:10
	AR_16x10
	// 16:9
	AR_16x9
	// 3:1
	AR_3x1
)

func (img *Image) AR(ar AspectRatio) (width, height uint) {
	width, height = img.WH()
	// TODO
	switch ar {
	case AR_1x1:
		if width > height {
			width = height
		} else {
			height = width
		}
	}
	return
}
