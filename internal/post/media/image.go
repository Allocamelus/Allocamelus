package media

import (
	"context"
	_ "embed"
	"errors"
	"mime/multipart"
	"os"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/pkg/imagedit"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jackc/pgx/v5"
)

const (
	SubPath = "posts/images"
)

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

	return nil
}

func checkCreateGetFile(img *imagedit.Image, imgHash string) (fileId int64, err error) {
	imgType := img.GetFormat()
	if !imgType.IsImage() {
		err = fileutil.ErrContentType
		return
	}
	ctx := context.Background()

	fileImagePath := fileutil.FilePath(selectorPath(imgHash, true))
	if fileutil.Exist(fileImagePath) {
		fileId, err = g.Data.Queries.GetPostMediaFileIDByHash(ctx, imgHash)
		// Not missing file in db
		if !errors.Is(err, pgx.ErrNoRows) {
			return
		}
	}

	// Check for reupload
	var dbHash string
	// Check for imgHash in db
	dbHash, err = g.Data.Queries.PostMediaHashCheck(ctx, imgHash)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return
		}
		dbHash = ""
	}

	// Missing file OR imgHash is newHash
	if dbHash != "" {
		// imgHash is newHash
		if imgHash != dbHash {
			if fileutil.Exist(fileutil.FilePath(selectorPath(dbHash, true))) {
				fileId, err = g.Data.Queries.GetPostMediaFileIDByHash(ctx, dbHash)
				// Not missing file in db
				if !errors.Is(err, pgx.ErrNoRows) {
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

	err = os.WriteFile(fileImagePath, imgOut, 0644)
	if err != nil {
		return
	}

	if dbHash != "" {
		fileId, err = g.Data.Queries.GetPostMediaFileIDByHash(ctx, dbHash)
		return
	}

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
