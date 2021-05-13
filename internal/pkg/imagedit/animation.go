package imagedit

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func isAnimation(mw *imagick.MagickWand) bool {
	return (mw.GetNumberImages() > 1)
}

type AnimationCallback func(img *Image) error

func (img *Image) IterateOver(callback AnimationCallback) error {
	if err := img.Check(); err != nil {
		return err
	}

	delay := img.MW.GetImageDelay()
	aImg, err := NewFromMW(img.MW.CoalesceImages())
	if err != nil {
		return err
	}
	defer aImg.Close()
	img.NewMW()

	img.MW.SetImageDelay(delay)
	frameCount := int(aImg.MW.GetNumberImages())

	for i := 0; i < frameCount; i++ {
		aImg.MW.SetIteratorIndex(i)
		aImgPart, err := NewFromMW(aImg.MW.GetImage())
		if err != nil {
			return err
		}
		err = callback(aImgPart)
		if err != nil {
			return err
		}
		img.MW.AddImage(aImgPart.MW)
		aImgPart.Close()
	}
	img.MW.ResetIterator()

	return nil
}
