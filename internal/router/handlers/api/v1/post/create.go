package post

import (
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

const (
	errUnauthorized      = "unauthorized"
	errInsufficientPerms = "insufficient-permissions"
)

type createRequest struct {
	Publish bool   `json:"publish" form:"publish"`
	Content string `json:"content" form:"content"`
}

type createResponce struct {
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

	newPost := post.New(sUser.UserID, request.Content, request.Publish)
	if err := newPost.ContentValid(); err != nil {
		return post403(c, err.Error())
	}

	if err := newPost.Insert(); logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, createResponce{
		Success: true,
		ID:      newPost.ID,
	})
}

func post403(c *fiber.Ctx, err string) error {
	return apierr.Err403(c, createResponce{Error: err})
}
