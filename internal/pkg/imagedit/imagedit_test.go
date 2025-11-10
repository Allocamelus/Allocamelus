package imagedit_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
)

const (
	a9s_jpg      = "./testFiles/a9s.jpg"
	a9s_test_jpg = "./testFiles/a9s.test.jpg"
)

var outDir = filepath.Join(os.TempDir(), "allocamelus-test")

func TestNewFromPath(t *testing.T) {
	img, err := imagedit.NewFromPath(a9s_jpg)
	if err != nil {
		t.Fatal(err)
	}
	defer img.Close()

	width, height := img.WH()
	if width <= 0 || height <= 0 {
		t.Fatal("Invalid width or height")
	}

	if img.GetFormat() != fileutil.JPG {
		t.Fatal("./a9s.jpg is not jpeg")
	}
}

func TestWriteToPath(t *testing.T) {
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

		err = img.WriteToPath(filepath.Join(outDir, "WriteToPath."+s))
		if err != nil {
			t.Fatal(err)
		}
	}
}
