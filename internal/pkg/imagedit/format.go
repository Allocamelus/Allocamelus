package imagedit

type Format string

const (
	WEBP Format = "WEBP"
	PNG  Format = "PNG"
	JPG  Format = "JPEG"
	GIF  Format = "GIF"
)

func (img *Image) GetFormat() Format {
	return Format(img.MW.GetImageFormat())
}
