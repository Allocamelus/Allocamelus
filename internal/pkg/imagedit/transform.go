package imagedit

import (
	"io/ioutil"
	"mime/multipart"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/blake2b"
)

func MPHtoImg(fileHead *multipart.FileHeader) (img *Image, b58hash string, err error) {
	file, err := fileHead.Open()
	if err != nil {
		return
	}
	defer file.Close()

	imageBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	img, err = NewFromBlob(imageBytes)
	if err != nil {
		return
	}

	rawHash := blake2b.Sum384(imageBytes)
	b58hash = base58.Encode(rawHash[:])

	return
}
