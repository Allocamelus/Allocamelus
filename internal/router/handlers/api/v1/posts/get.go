package posts

import (
	"strconv"

	"github.com/allocamelus/allocamelus/internal/pkg/dbutil"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type getResponse struct {
	*post.List
}

const perPage int64 = 15

// Get post handler
func Get(c *fiber.Ctx) error {
	page, _ := strconv.ParseInt(c.Query("p"), 10, 64)
	if page == 0 {
		page = 1
	}

	totalPosts, err := post.GetPublicTotal()
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, totalPosts)

	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	posts, err := post.GetPublicPosts(startNum, perPage)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	for _, p := range posts.Posts {
		p.MDtoHTMLContent()
	}

	return fiberutil.JSON(c, 200, getResponse{posts})
}
