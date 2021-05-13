package imagedit

import "github.com/allocamelus/allocamelus/internal/pkg/fileutil"

func (img *Image) Strip() error {
	if img.UseMW {
		if err := img.Check(); err != nil {
			return err
		}
		return img.MW.StripImage()
	}
	return nil
}

func (img *Image) Optimize() error {
	if img.UseMW {
		if err := img.Check(); err != nil {
			return err
		}
		if img.GetFormat() == fileutil.GIF && img.resized {
			img.NewMW(img.MW.OptimizeImageLayers())
		}
	}
	return nil
}
