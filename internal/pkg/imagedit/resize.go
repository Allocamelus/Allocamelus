package imagedit

import "gopkg.in/gographics/imagick.v3/imagick"

func (img *Image) Resize(width, height int) (err error) {
	if err := img.Check(); err != nil {
		return err
	}
	if img.UseMW {
		defer func(err error) {
			if err == nil {
				img.resized = true
			}
		}(err)

		callback := func(callbackImg *Image) error {
			return callbackImg.ResizeFilter(width, height, imagick.FILTER_LANCZOS2)
		}
		return img.IterateOver(callback)
	}

	return img.BlobToImg(img.Img.ForceResize(width, height))
}
func (img *Image) ResizeFilter(width, height int, filter imagick.FilterType) error {
	if err := img.Check(); err != nil {
		return err
	}
	return img.MW.ResizeImage(uint(width), uint(height), filter)
}
