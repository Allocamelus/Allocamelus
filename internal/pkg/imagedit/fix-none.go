//go:build !alpine
// +build !alpine

package imagedit

func (img *Image) Fix() {
	return
}
