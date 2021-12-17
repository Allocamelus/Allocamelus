package user

import (
	"github.com/allocamelus/allocamelus/internal/pkg/dbutil"
	"github.com/allocamelus/allocamelus/internal/post"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/shared"
	postsApi "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/posts"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

const perPage int64 = 15

//TODO: Share's Code with posts/get

// Posts posts handler
func Posts(c *fiber.Ctx) error {
	_, userID, hasErr, errApi := shared.GetUserNameIDResp(c)
	if hasErr {
		return errApi
	}

	page := fiberutil.ParamsInt64(c, "p")
	if page == 0 {
		page = 1
	}

	totalPosts, err := post.GetUserPostsTotal(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	startNum, totalPages := dbutil.ItemPageCalc(perPage, page, totalPosts)

	if page > totalPages && totalPages != 0 {
		return apierr.ErrNotFound(c)
	}

	posts, err := post.GetUserPosts(userID, startNum, perPage)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	// TODO Better Feed
	users := new(user.List)
	sessionUser := session.Context(c)
	for _, p := range posts.Posts {
		users.AddUser(sessionUser, p.UserID)
		p.MDtoHTMLContent()
	}

	return fiberutil.JSON(c, 200, postsApi.GetResponse{Posts: posts.Posts, Users: users.Users, Order: posts.Order})
}
