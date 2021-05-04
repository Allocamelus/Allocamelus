package follow

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

// requestsResp
type requestsResp struct {
	Requests map[int64]int64 `json:"requests"`
}

// Requests create follow
func Requests(c *fiber.Ctx) error {
	r, err := user.ListRequests(user.ContextSession(c).UserID)
	if err != nil {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, requestsResp{Requests: r})
}
