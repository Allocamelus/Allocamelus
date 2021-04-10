package imagedit

import (
	"errors"
)

func (img *Image) checkAnimation() {
	if err := img.Check(); err != nil {
		return
	}
	img.Animation = (img.MW.GetNumberImages() > 1)
}

type AnimationCallback func(img *Image) error

var ErrNotAnimation = errors.New("imagedit/animation: Error image not an animation")

func (img *Image) IterateOver(callback AnimationCallback) error {
	if err := img.Check(); err != nil {
		return err
	}
	if !img.Animation {
		return ErrNotAnimation
	}

	delay := img.MW.GetImageDelay()
	aImg, err := NewFromMW(img.MW.CoalesceImages())
	if err != nil {
		return err
	}
	defer aImg.Close()
	img.NewMW()

	var frameCount int
	// Only use first frame if TransformAnimation was false
	if img.TransformAnimation {
		img.MW.SetImageDelay(delay)
		frameCount = int(aImg.MW.GetNumberImages())
	} else {
		frameCount = 1
	}

	// deferred ln:16
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

	return nil
}
