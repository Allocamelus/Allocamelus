package imagedit

import (
	"errors"
	"io/ioutil"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/davidbyttow/govips/v2/vips"
)

type Image struct {
	img     *vips.ImageRef
	options *ImageOptions
}

type ImageOptions struct {
	FileType fileutil.Format
}

var (
	ErrBadType   = errors.New("imagedit: Error bad file type")
	ErrNilOutput = errors.New("imagedit: Error nil output")
)

func NewFromPath(imagePath string) (*Image, error) {
	imgBlob, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	return NewFromBlob(imgBlob)
}

func NewFromBlob(blob []byte) (*Image, error) {
	img := new(Image)
	return img, img.BlobToImg(blob)
}

// Decodes blob &
func (img *Image) BlobToImg(blob []byte) (err error) {
	importParams := vips.NewImportParams()
	// Set to support animation
	importParams.NumPages.Set(-1)

	img.img, err = vips.LoadImageFromBuffer(blob, importParams)
	if err != nil {
		return
	}

	// SetPages to one if not animation
	if img.Pages() == 1 {
		img.img.SetPages(1)
	}

	// Get image format
	imgFmt := fileutil.ExtensionToFormat(img.img.Format().FileExt())
	if !imgFmt.IsImage() {
		return ErrBadType
	}

	img.options = &ImageOptions{
		FileType: imgFmt,
	}
	return
}

func (img *Image) Export() ([]byte, error) {
	var err error
	var outputImg []byte

	img.Fix()

	switch img.options.FileType {
	case fileutil.GIF:
		ep := vips.NewGifExportParams()
		ep.Quality = 90
		ep.StripMetadata = true
		outputImg, _, err = img.img.ExportGIF(ep)
	case fileutil.JPG:
		ep := vips.NewJpegExportParams()
		ep.Quality = 90
		ep.StripMetadata = true
		outputImg, _, err = img.img.ExportJpeg(ep)
	case fileutil.PNG:
		ep := vips.NewPngExportParams()
		ep.Compression = 9
		ep.StripMetadata = true
		outputImg, _, err = img.img.ExportPng(ep)
	case fileutil.WEBP:
		ep := vips.NewWebpExportParams()
		ep.Quality = 90
		ep.StripMetadata = true
		outputImg, _, err = img.img.ExportWebp(ep)
	}
	if err != nil {
		return nil, err
	}
	if outputImg == nil {
		return nil, ErrNilOutput
	}

	return outputImg, nil
}

func (img *Image) WriteToPath(imagePath string) error {
	outputImg, err := img.Export()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(imagePath, outputImg, 0644)
}

func (img *Image) Close() {
	img.img.Close()
}

// WH returns image width & height
//
//	return width, height int, err error
func (img *Image) WH() (width, height int) {
	width = img.img.Width()
	if img.Pages() == 1 {
		height = img.img.Height()
	} else {
		height = img.img.PageHeight()
	}
	return
}

// Pages returns number of pages
func (img *Image) Pages() (pages int) {
	return img.img.Pages()
}
