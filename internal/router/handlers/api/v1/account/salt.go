package account

import (
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user/auth"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type SaltRequest struct {
	UserName string `json:"userName" form:"userName"`
}

type SaltResponse struct {
	Salt  string `json:"salt"`
	Error string `json:"error,omitempty"`
}

func (s *SaltResponse) error(c *fiber.Ctx, err string) error {
	s.Error = err
	return apierr.Err422(c, s)
}

// Salt get user's salt
func Salt(c *fiber.Ctx) error {
	request := new(SaltRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	userID, err := auth.CanLogin(request.UserName)
	if err != nil {
		switch err {
		case auth.ErrInvalidUsername:
			return new(SaltResponse).error(c, errInvalidUsernamePassword)
		case auth.ErrUnverifiedEmail:
			return new(SaltResponse).error(c, errUnverifiedEmail)
		}
		logger.Error(err)
		return apierr.ErrSomethingWentWrong(c)
	}

	salt, err := key.GetSalt(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, SaltResponse{Salt: salt})
}
