package imagedit

import (
	"io"
	"mime/multipart"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/blake2b"
)

func MPHtoImg(fileHead *multipart.FileHeader) (img *Image, b58hash string, err error) {
	file, err := fileHead.Open()
	if err != nil {
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	img, err = NewFromBlob(imageBytes)
	if err != nil {
		return
	}

	b58hash = HashEncode(imageBytes)

	return
}

func HashEncode(bytes []byte) string {
	rawHash := blake2b.Sum384(bytes)
	return base58.Encode(rawHash[:])
}

// transformPages apply modifier to page(s)
func (img *Image) transformPages(modifier func(img *vips.ImageRef) error) error {
	pages := img.Pages()

	// Skip unneeded steps
	if pages == 1 {
		return modifier(img.img)
	}

	// Get metadata before flattening
	width := img.img.Width()
	pageHeight := img.img.PageHeight()
	pageDelay, err := img.img.PageDelay()
	if err != nil {
		return err
	}

	// flatten image
	img.img.SetPages(1)
	img.img.SetPageHeight(img.img.Height())

	imgPages, err := img.img.Copy()
	if err != nil {
		return err
	}

	// Extract as first page
	if err = imgPages.ExtractArea(0, 0, width, pageHeight); err != nil {
		return err
	}

	if err = modifier(imgPages); err != nil {
		return err
	}

	newPageHeight := imgPages.Height()

	for i := 1; i < pages; i++ {
		page, err := img.img.Copy()
		if err != nil {
			return err
		}
		defer page.Close()

		// Extract page i
		if err = page.ExtractArea(0, pageHeight*i, width, pageHeight); err != nil {
			return err
		}

		if err = modifier(page); err != nil {
			return err
		}

		// Append to

		if err = imgPages.Join(page, vips.DirectionVertical); err != nil {
			return err
		}
	}

	if err = imgPages.SetPages(pages); err != nil {
		return err
	}

	if err = imgPages.SetPageHeight(newPageHeight); err != nil {
		return err
	}

	if err = imgPages.SetPageDelay(pageDelay); err != nil {
		return err
	}

	img.img.Close()
	img.img = *&imgPages
	return nil
}
