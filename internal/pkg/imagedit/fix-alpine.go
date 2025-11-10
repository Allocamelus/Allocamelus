//go:build alpine
// +build alpine

package imagedit

import "github.com/allocamelus/allocamelus/internal/pkg/fileutil"

// Fix fixes VipsOperation: class "gifsave_buffer" not found on alpine
// by disabling gif saving
func (img *Image) Fix() {
	if img.options.FileType == fileutil.GIF {
		img.options.FileType = fileutil.WEBP
	}
}
