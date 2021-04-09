package update

import (
	"log"
	"os"
	"path/filepath"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/avatar"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"github.com/gofiber/fiber/v2"
)

type AvatarResp struct {
	Success   bool   `json:"success"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Error     string `json:"error,omitempty"`
}

// Avatar Update handler
func Avatar(c *fiber.Ctx) error {
	// User can't post images
	if !user.ContextSession(c).Perms.CanUploadMedia() {
		return apierr.ErrUnauthorized403(c)
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	log.Println(file.Filename, file.Size)

	if err := avatar.ValidateMpFileHeader(file); err != nil {
		if err == avatar.ErrSomethingWentWrong {
			return apierr.ErrSomethingWentWrong(c)
		}
		return apierr.Err422(c, AvatarResp{Error: err.Error()})
	}

	imgDir := dirutil.RandomTmpDir()
	// Remove image tmp dir and all it's children
	defer os.RemoveAll(imgDir)
	imgPath := filepath.Join(imgDir, random.StringBase64(8))
	// Save file to tmp dir
	if err := c.SaveFile(file, imgPath); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	newUrl, err := avatar.TransformAndSave(user.ContextSession(c).UserID, imgPath)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, AvatarResp{Success: true, AvatarUrl: newUrl})
}

// RemoveAvatar handler
func RemoveAvatar(c *fiber.Ctx) error {
	userId := user.ContextSession(c).UserID

	hasAvatar, err := avatar.HasAvatar(userId)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	if hasAvatar {
		if err := avatar.Remove(userId); logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	return c.SendStatus(204)
}
