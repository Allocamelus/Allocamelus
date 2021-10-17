package imagedit

import (
	"errors"
	"io/ioutil"

	"github.com/discord/lilliput"
)

type Image struct {
	Img     lilliput.Decoder
	options *lilliput.ImageOptions
}

// EncodeOptions default encoding options
var EncodeOptions = map[string]map[int]int{
	".jpg":  map[int]int{lilliput.JpegQuality: 90},
	".jpeg": map[int]int{lilliput.JpegQuality: 90},
	".png":  map[int]int{lilliput.PngCompression: 9},
	".webp": map[int]int{lilliput.WebpQuality: 90},
}

var (
	ErrBadType = errors.New("imagedit: Error bad file type")
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
	img.Img, err = lilliput.NewDecoder(blob)
	if err != nil {
		return
	}

	// Get image format
	imgFmt := strToFmt(img.Img.Description())
	if !imgFmt.IsImage() {
		return ErrBadType
	}
	// Get image header
	header, err := img.Img.Header()
	if err != nil {
		return
	}

	img.options = &lilliput.ImageOptions{
		FileType:             imgFmt.FileExt(),
		Width:                header.Width(),
		Height:               header.Height(),
		ResizeMethod:         lilliput.ImageOpsNoResize,
		NormalizeOrientation: true,
		EncodeOptions:        EncodeOptions[imgFmt.FileExt()],
	}
	return
}

func (img *Image) WriteToPath(imagePath string) (err error) {
	// get ready to resize image,
	// using 8192x8192 maximum resize buffer size
	ops := lilliput.NewImageOps(8192)
	defer ops.Close()

	// create a buffer to store the output image, 50MB in this case
	outputImg := make([]byte, 50*1024*1024)
	// resize and transcode image
	outputImg, err = ops.Transform(img.Img, img.options, outputImg)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(imagePath, outputImg, 0644)
	return
}

func (img *Image) Close() {
	img.Img.Close()
}

// WH returns image width & height
//  return width, height int, err error
func (img *Image) WH() (width, height int) {
	return img.options.Width, img.options.Height
}
