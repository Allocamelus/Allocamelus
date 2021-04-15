package post

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/allocamelus/allocamelus/internal/pkg/dirutil"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"github.com/gofiber/fiber/v2"
)

const (
	errUnauthorized      = "unauthorized"
	errInsufficientPerms = "insufficient-permissions"
)

type createRequest struct {
	Publish bool                    `json:"publish" form:"publish"`
	Content string                  `json:"content" form:"content"`
	Images  []*multipart.FileHeader `form:"images[]"`
}

type createResponse struct {
	Success bool   `json:"success"`
	ID      int64  `json:"id,omitempty"`
	Error   string `json:"error,omitempty"`
}

// TODO: Rate limiting

// Create post handler
func Create(c *fiber.Ctx) error {
	sUser := user.ContextSession(c)
	if !sUser.Perms.CanPost() {
		return post403(c, errInsufficientPerms)
	}

	request := new(createRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.Images = form.File["images[]"]

	newPost := post.New(sUser.UserID, request.Content, request.Publish)
	if err := newPost.ContentValid(); err != nil {
		return post403(c, err.Error())
	}
	if err := newPost.Insert(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	for _, v := range request.Images {
		if err := media.ValidateMpFileHeader(v); err != nil {
			if err == fileutil.ErrSomethingWentWrong {
				return apierr.ErrSomethingWentWrong(c)
			}
			return apierr.Err422(c, createResponse{Error: err.Error()})
		}
		imgDir := dirutil.RandomTmpDir()
		// Remove image tmp dir and all it's children
		defer os.RemoveAll(imgDir)
		imgPath := filepath.Join(imgDir, random.StringBase64(8))
		// Save file to tmp dir
		if err := c.SaveFile(v, imgPath); logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
		err := media.TransformAndSave(newPost.ID, imgPath)
		if logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	return fiberutil.JSON(c, 200, createResponse{
		Success: true,
		ID:      newPost.ID,
	})
}

func post403(c *fiber.Ctx, err string) error {
	return apierr.Err403(c, createResponse{Error: err})
}
