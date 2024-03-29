package update

import (
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user/avatar"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
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
	if !session.Context(c).Perms.CanUploadMedia() {
		return apierr.ErrUnauthorized403(c)
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	if err := avatar.ValidateMpFileHeader(file); err != nil {
		if err == fileutil.ErrSomethingWentWrong {
			return apierr.ErrSomethingWentWrong(c)
		}
		return apierr.Err422(c, AvatarResp{Error: err.Error()})
	}

	newUrl, err := avatar.TransformAndSave(session.Context(c).UserID, file)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, AvatarResp{Success: true, AvatarUrl: newUrl})
}

// RemoveAvatar handler
func RemoveAvatar(c *fiber.Ctx) error {
	userId := session.Context(c).UserID

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
