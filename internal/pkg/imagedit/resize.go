package imagedit

import (
	"github.com/davidbyttow/govips/v2/vips"
)

// Resize image
func (img *Image) Resize(width, height int) error {
	oldWidth, oldHeight := img.WH()
	hScale := float64(width) / float64(oldWidth)
	vScale := float64(height) / float64(oldHeight)
	return img.transformPages(func(img *vips.ImageRef) error {
		return img.ResizeWithVScale(hScale, vScale, vips.KernelLanczos3)
	})
}

// Crop image
func (img *Image) Crop(width, height int) error {
	return img.transformPages(func(img *vips.ImageRef) error {
		return img.SmartCrop(width, height, vips.InterestingCentre)
	})
}

// CropAR crops image using aspect ratio
func (img *Image) CropAR(ar AspectRatio) error {
	width, height := img.AR(ar)
	return img.Crop(width, height)
}

// Thumbnail resizes and crops image
func (img *Image) Thumbnail(width, height int) error {
	return img.transformPages(func(img *vips.ImageRef) error {
		return img.Thumbnail(width, height, vips.InterestingCentre)
	})
}

// ThumbnailAR resizes and crops image using aspect ratio
func (img *Image) ThumbnailAR(ar AspectRatio) error {
	width, height := img.AR(ar)
	return img.Thumbnail(width, height)
}
