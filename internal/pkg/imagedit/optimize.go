package imagedit

func (img *Image) Strip() error {
	if err := img.Check(); err != nil {
		return err
	}
	return img.MW.StripImage()
}

func (img *Image) Optimize() error {
	if err := img.Check(); err != nil {
		return err
	}
	if img.MW.GetImageFormat() == GIF && img.resized {
		img.NewMW(img.MW.OptimizeImageLayers())
	}
	return nil
}
