package post

import (
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type getResponse struct {
	Post *post.Post `json:"post"`
	User *user.User `json:"user"`
}

// Get post handler
func Get(c *fiber.Ctx) error {
	p, errFunc := getForUser(c)
	if errFunc != nil {
		return errFunc(c)
	}

	p.MDtoHTMLContent()
	u, err := user.GetPublic(p.UserID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	return fiberutil.JSON(c, 200, getResponse{Post: p, User: &u})
}

func getForUser(c *fiber.Ctx) (*post.Post, fiber.Handler) {
	postID := fiberutil.ParamsInt64(c, "id")
	if postID == 0 {
		return nil, apierr.ErrNotFound
	}

	p, err := post.GetForUser(int64(postID), user.ContextSession(c).UserID)
	if err != nil {
		if err != post.ErrNoPost {
			logger.Error(err)
			return nil, apierr.ErrSomethingWentWrong
		}
		return nil, apierr.ErrNotFound
	}
	return p, nil
}
