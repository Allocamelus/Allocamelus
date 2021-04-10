package imagedit

import "gopkg.in/gographics/imagick.v3/imagick"

func (img *Image) Resize(width, height uint) error {
	return img.ResizeFilter(width, height, imagick.FILTER_LANCZOS2)
}

func (img *Image) ResizeFilter(width, height uint, filter imagick.FilterType) error {
	if err := img.Check(); err != nil {
		return err
	}

	if img.Animation {
		callback := func(callbackImg *Image) error {
			return callbackImg.ResizeFilter(width, height, filter)
		}
		return img.IterateOver(callback)
	}

	return img.MW.ResizeImage(width, height, filter)
}
