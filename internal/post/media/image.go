package media

import (
	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
)

const MaxHightWidth uint = 7680

func TransformAndSave(postID int64, tmpImagePath string) (err error) {
	img, err := imagedit.NewFromPath(tmpImagePath)
	if err != nil {
		return
	}
	defer img.Close()

	err = img.Strip()
	if err != nil {
		return
	}
	// Allow Animations
	img.TransformAnimation = true
	width, height := img.WH()
	if width > MaxHightWidth || height > MaxHightWidth || img.Animation {
		newWidth, newHeight := img.ARMaxSize(imagedit.AR_Image, MaxHightWidth)
		err = img.Resize(newWidth, newHeight)
		if err != nil {
			return
		}
	}
	err = img.Optimize()
	if err != nil {
		return
	}

	selector := random.StringBase58(16)

	mediaId, err := Insert(postID, Image, selector)
	if err != nil {
		return
	}

	fileImagePath := fileutil.FilePath(selectorPath(mediaId, selector, Image))

	logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(mediaId, "", Image))))

	err = img.WriteToPath(fileImagePath)

	if err != nil {
		return
	}

	return
}
