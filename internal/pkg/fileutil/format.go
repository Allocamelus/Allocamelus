//go:generate msgp

package fileutil

type Format int

const (
	NONE = Format(iota)
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
		return ".webp"
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

func ContentTypeToFormat(ct string) Format {
	switch ct {
	case "image/gif":
		return GIF
	case "image/jpeg":
		return JPG
	case "image/png":
		return PNG
	case "image/webp":
		return WEBP
	}
	return NONE
}

func ExtensionToFormat(ext string) Format {
	switch ext {
	case ".gif":
		return GIF
	case ".jpeg", ".jpg":
		return JPG
	case ".png":
		return PNG
	case ".webp":
		return WEBP
	}
	return NONE
}
