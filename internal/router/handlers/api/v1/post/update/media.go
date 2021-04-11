package update

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

type MediaResp struct {
	Success   bool   `json:"success"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Error     string `json:"error,omitempty"`
}

// Avatar Update handler
func Media(c *fiber.Ctx) error {
	// User can't post images
	if !user.ContextSession(c).Perms.CanUploadMedia() {
		return apierr.ErrUnauthorized403(c)
	}
	return fiberutil.JSON(c, 200, MediaResp{Success: true})
}

// TODO
// RemoveMedia handler
func RemoveMedia(c *fiber.Ctx) error {
	return c.SendStatus(204)
}
