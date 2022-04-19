package imagedit_test

import (
	_ "embed"
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
)

var (
	//go:embed testFiles/DNA.gif
	dnaGIF []byte
	//go:embed testFiles/a9s.jpg
	a9sJPG []byte
	//go:embed testFiles/a9s.png
	a9sPNG []byte
	//go:embed testFiles/a9s.webp
	a9sWEBP []byte
)

type image struct {
	Img []byte
	Fmt fileutil.Format
}

var strToFmtList = map[string]*image{
	"DNA.gif":  {Img: dnaGIF, Fmt: fileutil.GIF},
	"a9s.jpg":  {Img: a9sJPG, Fmt: fileutil.JPG},
	"a9s.png":  {Img: a9sPNG, Fmt: fileutil.PNG},
	"a9s.webp": {Img: a9sWEBP, Fmt: fileutil.WEBP},
}

func TestFormat(t *testing.T) {
	for _, i := range strToFmtList {
		img, err := imagedit.NewFromBlob(i.Img)
		if err != nil {
			t.Fatal(err)
		}
		defer img.Close()
		if img.GetFormat() != i.Fmt {
			t.Fatal("Incorrect format")
		}
	}
}
