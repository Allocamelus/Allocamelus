package media

import (
	"database/sql"
	_ "embed"
	"io/ioutil"
	"mime/multipart"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

const (
	SubPath = "posts/images"
)

var (
	//go:embed sql/get/hashCheck.sql
	qGetHashCheck   string
	preGetHashCheck *sql.Stmt
	//go:embed sql/get/fileId.sql
	qGetFileId   string
	preGetFileId *sql.Stmt
)

func init() {
	data.PrepareQueuer.Add(&preGetHashCheck, qGetHashCheck)
	data.PrepareQueuer.Add(&preGetFileId, qGetFileId)
	//data.PrepareQueuer.Add(&preInsert, qInsert)
}

func TransformAndSave(postID int64, imageMPH *multipart.FileHeader, alt string) error {
	img, b58hash, err := imagedit.MPHtoImg(imageMPH)
	if err != nil {
		return err
	}
	defer img.Close()

	fileId, err := checkCreateGetFile(img, b58hash)
	if err != nil {
		return err
	}

	err = Insert(postID, alt, fileId)
	if err != nil {
		return err
	}

	return err
}

func checkCreateGetFile(img *imagedit.Image, imgHash string) (fileId int64, err error) {
	imgType := img.GetFormat()
	if !imgType.IsImage() {
		err = fileutil.ErrContentType
		return
	}

	fileImagePath := fileutil.FilePath(selectorPath(imgHash, true))
	if fileutil.Exist(fileImagePath) {
		err = preGetFileId.QueryRow(imgHash).Scan(&fileId)
		// Not missing file in db
		if err != sql.ErrNoRows {
			return
		}
	}

	// Check for reupload
	var dbHash string
	// Check for imgHash in db
	err = preGetHashCheck.QueryRow(imgHash, imgHash).Scan(&dbHash)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	// Missing file OR imgHash is newHash
	if dbHash != "" {
		// imgHash is newHash
		if imgHash != dbHash {
			if fileutil.Exist(fileutil.FilePath(selectorPath(dbHash, true))) {
				err = preGetFileId.QueryRow(dbHash).Scan(&fileId)
				// Not missing file in db
				if err != sql.ErrNoRows {
					return
				}
			}
			// Missing dbHash file
			imgHash = dbHash
			fileImagePath = fileutil.FilePath(selectorPath(imgHash, true))
		}
	}

	// Make folders for image
	logger.Error(dirutil.MakeDir(fileutil.FilePath(selectorPath(imgHash, false))))

	imgOut, err := img.Export()
	if err != nil {
		return
	}

	err = ioutil.WriteFile(fileImagePath, imgOut, 0644)
	if err != nil {
		return
	}

	if dbHash == "" {
		newHash := imagedit.HashEncode(imgOut)
		width, height := img.WH()
		fileId, err = InsertFile(
			&Media{
				FileType: imgType,
				Meta: &Meta{
					Width:  int64(width),
					Height: int64(height),
				}},
			imgHash,
			newHash,
		)
		return
	}

	err = preGetFileId.QueryRow(imgHash).Scan(&fileId)
	return
}
