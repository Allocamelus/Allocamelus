package imagedit_test

import (
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
)

var strToFmtList = map[string]fileutil.Format{
	"./testFiles/DNA.gif":  fileutil.GIF,
	a9s_jpg:                fileutil.JPG,
	"./testFiles/a9s.png":  fileutil.PNG,
	"./testFiles/a9s.webp": fileutil.WEBP,
}

func TestFormat(t *testing.T) {
	for s, f := range strToFmtList {
		img, err := imagedit.NewFromPath(s)
		if err != nil {
			t.Fatal(err)
		}
		defer img.Close()
		if img.GetFormat() != f {
			t.Fatal("Incorrect format")
		}
	}
}
