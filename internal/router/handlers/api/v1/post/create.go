package post

import (
	"mime/multipart"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/pkg/errtools"
	"github.com/allocamelus/allocamelus/internal/pkg/fileutil"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/post/media"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	Publish   bool                    `json:"publish" form:"publish"`
	Content   string                  `json:"content" form:"content"`
	Images    []*multipart.FileHeader `form:"images[]"`
	ImageAlts []string                `form:"imageAlts[]"`
}

// TODO: Rate limiting

// Create post handler
func Create(c *fiber.Ctx) error {
	sUser := user.ContextSession(c)
	if !sUser.Perms.CanPost() {
		return post403(c, errtools.ErrInsufficientPerms.Error())
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

	imageAltLen := len(request.ImageAlts)

	for k, v := range request.Images {
		if err := media.ValidateMpFileHeader(v); err != nil {
			if err == fileutil.ErrSomethingWentWrong {
				return apierr.ErrSomethingWentWrong(c)
			}
			return apierr.Err422(c, shared.SuccessErrResp{Error: err.Error()})
		}
		var alt string
		if imageAltLen > k {
			// Truncate alt to 512
			altLen := len(request.ImageAlts[k])
			if altLen > 512 {
				altLen = 512
			}
			alt = request.ImageAlts[k][:altLen]
		} else {
			alt = "Image #" + strconv.Itoa(k+1) + " For Post:" + strconv.Itoa(int(newPost.ID))
		}
		err := media.TransformAndSave(newPost.ID, v, alt)
		if logger.Error(err) {
			return apierr.ErrSomethingWentWrong(c)
		}
	}

	return fiberutil.JSON(c, 200, shared.SuccessIDErrResp{
		Success: true,
		ID:      newPost.ID,
	})
}

func post403(c *fiber.Ctx, err string) error {
	return apierr.Err403(c, shared.SuccessErrResp{Error: err})
}
