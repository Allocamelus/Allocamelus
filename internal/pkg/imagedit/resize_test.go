package imagedit_test

import (
	"path/filepath"
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
)

func TestResize(t *testing.T) {
	err := dirutil.MakeDir(outDir)
	if err != nil {
		t.Fatal(err)
	}
	for s, i := range strToFmtList {
		img, err := imagedit.NewFromBlob(i.Img)
		if err != nil {
			t.Fatal(err)
		}
		defer img.Close()

		width, height := img.WH()
		err = img.Resize(width/2, height/2)
		if err != nil {
			t.Fatal(err)
		}
		nWidth, nHeight := img.WH()
		if nWidth != width/2 {
			t.Fatal("Incorrect width")
		}
		if nHeight != height/2 {
			t.Fatal("Incorrect height")
		}

		err = img.WriteToPath(filepath.Join(outDir, "Resize."+s))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestThumbnailAR(t *testing.T) {
	err := dirutil.MakeDir(outDir)
	if err != nil {
		t.Fatal(err)
	}
	for s, i := range strToFmtList {
		img, err := imagedit.NewFromBlob(i.Img)
		if err != nil {
			t.Fatal(err)
		}
		defer img.Close()

		err = img.ThumbnailAR(imagedit.AR_1x1)
		if err != nil {
			t.Fatal(err)
		}

		if width, height := img.WH(); width != height {
			t.Fatal("Incorrect AR")
		}
		err = img.WriteToPath(filepath.Join(outDir, "Thumbnail."+s))
		if err != nil {
			t.Fatal(err)
		}
	}
}
