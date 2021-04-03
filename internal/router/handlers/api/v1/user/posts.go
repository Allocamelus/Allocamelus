package user

import (
	"strconv"

	"github.com/allocamelus/allocamelus/internal/pkg/dbutil"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	postsApi "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/posts"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

const perPage int64 = 15

//TODO: Share's Code with posts/get

// Posts posts handler
func Posts(c *fiber.Ctx) error {
	_, userID, hasErr, errApi := getUserNameID(c)
	if hasErr {
		return errApi
	}

	page, _ := strconv.ParseInt(c.Query("p"), 10, 64)
	if page == 0 {
		page = 1
	}

	totalPosts, err := post.GetPublicUserTotal(userID)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, totalPosts)

	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	posts, err := post.GetPublicUserPosts(userID, startNum, perPage)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	// TODO Better Feed
	users := new(user.List)
	for _, p := range posts.Posts {
		users.AddUser(p.UserID)
		p.MDtoHTMLContent()
	}

	return fiberutil.JSON(c, 200, postsApi.GetResponse{Posts: posts.Posts, Users: users.Users, Order: posts.Order})
}