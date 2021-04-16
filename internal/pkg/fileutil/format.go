//go:generate msgp

package fileutil

type Format int

const (
	NONE = iota
	GIF
	JPG
	PNG
	WEBP
)

func (f Format) FileExt() string {
	switch f {
	case GIF:
		return ".gif"
	case JPG:
		return ".jpg"
	case PNG:
		return ".png"
	case WEBP:
		return ".web"
	}
	return ""
}

func (f Format) IsImage() bool {
	switch f {
	case GIF, JPG, PNG, WEBP:
		return true
	}
	return false
}
