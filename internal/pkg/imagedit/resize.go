package imagedit

import (
	"github.com/discord/lilliput"
)

func (img *Image) Resize(width, height int) {
	img.options.Width = width
	img.options.Height = height
	img.options.ResizeMethod = lilliput.ImageOpsResize
}
