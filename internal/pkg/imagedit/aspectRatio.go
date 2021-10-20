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

func (img *Image) AR(ar AspectRatio) (width, height int) {
	width, height = img.WH()
	// TODO
	if ar == AR_1x1 {
		if width > height {
			width = height
		} else {
			height = width
		}
	}
	return
}

func (img *Image) ARMaxSize(ar AspectRatio, maxPx int) (width, height int) {
	width, height = img.AR(ar)

	maxPxF := float64(maxPx)
	// Resize for width
	wf, hf := resizeAr(float64(width), float64(height), maxPxF)
	// resize for height
	hf, wf = resizeAr(hf, wf, maxPxF)

	width = int(wf)
	height = int(hf)
	return
}

// resizeAr reduces x to maxX if x > maxX
func resizeAr(x, y, maxX float64) (float64, float64) {
	if x > maxX {
		resizeFloat := maxX / x
		x = resizeFloat * x
		y = resizeFloat * y
	}
	return x, y
}
