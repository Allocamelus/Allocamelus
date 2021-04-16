package media

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/blake2b"
)

const MaxHightWidth uint = 7680

func TransformAndSave(postID int64, imageMPH *multipart.FileHeader, alt string) (err error) {
	file, err := imageMPH.Open()
	if err != nil {
		return
	}
	defer file.Close()

	imageBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	img, err := imagedit.NewFromBlob(imageBytes)
	if err != nil {
		return
	}
	defer img.Close()

	imgType := ImageditFmtToType(img.GetFormat())
	if imgType == None {
		err = fileutil.ErrContentType
		return
	}

	imgHash := blake2b.Sum384(imageBytes)
	encoded := base58.Encode(imgHash[:])

	fileImagePath := fileutil.FilePath(selectorPath(encoded, imgType, true))

	_, err = os.Stat(fileImagePath)
	// Check for image for deduplication
	if errors.Is(err, os.ErrNotExist) {
		err = img.Strip()
		if err != nil {
			return
		}
		// Allow Animations
		img.TransformAnimation = true
		width, height := img.WH()
		var newWidth, newHeight uint
		if width > MaxHightWidth || height > MaxHightWidth || img.Animation {
			newWidth, newHeight = img.ARMaxSize(imagedit.AR_Image, MaxHightWidth)
		} else {
			newWidth = width
			newHeight = height
		}
		// Resize to prevent non images
		err = img.Resize(newWidth, newHeight)
		if err != nil {
			return
		}
		err = img.Optimize()
		if err != nil {
			return
		}
	} else {
		logger.Error(err)
	}

	width, height := img.WH()

	err = Insert(postID, Media{MediaType: imgType, Meta: Meta{Alt: alt, Width: int64(width), Height: int64(height)}}, encoded)
	if err != nil {
		return
	}

	logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(encoded, imgType, false))))

	err = img.WriteToPath(fileImagePath)
	if err != nil {
		return
	}

	return
}
