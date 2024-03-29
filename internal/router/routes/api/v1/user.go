package v1

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user"
	emailtoken "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/email-token"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/follow"
	userupdate "github.com/allocamelus/allocamelus/internal/router/handlers/api/v1/user/update"
	"github.com/allocamelus/allocamelus/internal/router/middleware"
	"github.com/gofiber/fiber/v2"
)

// User routes
func User(api fiber.Router) {
	// /api/v1/user
	u := api.Group("/user")

	// /api/v1/user/email-token
	uET := u.Group("/email-token")
	uET.Post("/", emailtoken.Create)
	// /api/v1/user/email-token/validate
	uET.Post("/validate", emailtoken.Validate)

	// /api/v1/user/follow
	uF := u.Group("/follow", middleware.Protected)
	uF.Get("/requests", follow.Requests)

	/* TODO: // /api/v1/user/password-reset
	uPR := u.Group("/password-reset")
	// /api/v1/user/password-reset/token
	uPR.Post("/token", passreset.CreateToken)
	// /api/v1/user/password-reset/validate
	uPRV := uPR.Group("/validate")
	// /api/v1/user/password-reset/validate/token
	uPRV.Post("/token", passresetval.Token)*/

	// /api/v1/user/:userName
	uUN := u.Group("/:userName")
	uUN.Get("/", user.Get)
	// /api/v1/user/:userName/posts
	uUN.Get("/posts", middleware.ProtectedPubOrFollow, user.Posts)
	// TODO: require revocation of cert or something
	// /api/v1/user/:userName/delete
	uUN.Delete("/delete",
		middleware.Protected,
		middleware.ProtectedSelfOnly,
		user.Delete,
	)
	userUpdate(uUN)
	userFollow(uUN)
}

func userUpdate(un fiber.Router) {
	// /api/v1/user/:userName/update
	unGroup := un.Group("/update",
		middleware.Protected,
		middleware.ProtectedSelfOnly,
	)
	// /api/v1/user/:userName/update/avatar
	unGroup.Post("/avatar",
		userupdate.Avatar,
	)
	unGroup.Delete("/avatar",
		userupdate.RemoveAvatar,
	)
	// /api/v1/user/:userName/update/bio
	unGroup.Post("/bio",
		userupdate.Bio,
	)
	// /api/v1/user/:userName/update/name
	unGroup.Post("/name",
		userupdate.Name,
	)
	// /api/v1/user/:userName/update/type
	unGroup.Post("/type",
		userupdate.Type,
	)
}

func userFollow(un fiber.Router) {
	// /api/v1/user/:userName/follow
	unFollow := un.Group("/follow", middleware.Protected)
	// Followers
	unFollow.Post("/", follow.Post)
	unFollow.Delete("/", follow.Delete)
	// Friends
	// /api/v1/user/:userName/follow/accept
	unFollow.Post("/accept", follow.Accept)
	// /api/v1/user/:userName/follow/decline
	unFollow.Delete("/decline", follow.Delete)
}
