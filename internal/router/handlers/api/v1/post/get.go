package post

import (
	"strconv"

	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type getResponse struct {
	post.Post
}

// Get post handler
func Get(c *fiber.Ctx) error {
	postIDstr := c.Params("id")
	if len(postIDstr) == 0 {
		return apierr.ErrNotFound(c)
	}
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		return apierr.ErrNotFound(c)
	}

	p, err := post.GetForUser(int64(postID), user.ContextSession(c))
	if err != nil {
		if err != post.ErrNoPost {
			logger.Error(err)
			return apierr.ErrSomthingWentWrong(c)
		}
		return apierr.ErrNotFound(c)
	}

	p.MDtoHTMLContent()
	return fiberutil.JSON(c, 200, getResponse{p})
}
