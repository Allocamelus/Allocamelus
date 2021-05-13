package posts

import (
	"strconv"

	"github.com/allocamelus/allocamelus/internal/pkg/dbutil"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// GetResponse posts response
type GetResponse struct {
	Posts post.ListPosts  `json:"posts"`
	Users user.ListUsers  `json:"users"`
	Order map[int64]int64 `json:"order"`
}

const perPage int64 = 15

// Get posts handler
func Get(c *fiber.Ctx) error {
	page, _ := strconv.ParseInt(c.Query("p"), 10, 64)
	if page == 0 {
		page = 1
	}

	totalPosts, err := post.GetPublicTotal(user.ContextSession(c))
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, totalPosts)

	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	posts, err := post.GetPublicPosts(startNum, perPage, user.ContextSession(c))
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	users := new(user.List)
	sessionUser := user.ContextSession(c)
	for _, p := range posts.Posts {
		users.AddUser(sessionUser, p.UserID)
		p.MDtoHTMLContent()
	}

	return fiberutil.JSON(c, 200, GetResponse{Posts: posts.Posts, Users: users.Users, Order: posts.Order})
}
