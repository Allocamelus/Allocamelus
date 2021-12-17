package account

import (
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/auth"
	"github.com/allocamelus/allocamelus/internal/user/session"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// AuthRequest struct
type AuthRequest struct {
	UserName string `json:"userName"`
	AuthKey  string `json:"authkey"`
	Remember bool   `json:"remember"`
	Captcha  string `json:"captcha"`
}

func (t *AuthRequest) trimSpace() {
	t.UserName = strings.TrimSpace(t.UserName)
	t.AuthKey = strings.TrimSpace(t.AuthKey)
	t.Captcha = strings.TrimSpace(t.Captcha)
}

// AuthResponse struct
type AuthResponse struct {
	Success        bool           `json:"success"`
	PrivateArmored pgp.PrivateKey `json:"privateArmored"`
	User           user.User      `json:"user,omitempty"`
	Error          string         `json:"error,omitempty"`
	// Require Captcha
	Captcha string `json:"captcha,omitempty"`
}

func (s *AuthResponse) error(c *fiber.Ctx, err string, captcha ...string) error {
	s.Error = err
	if len(captcha) != 0 {
		s.Captcha = captcha[0]
	}
	return apierr.Err422(c, s)
}

const (
	errInvalidCaptcha          = "invalid-captcha"
	errInvalidUsernamePassword = "invalid-username-password"
	errUnverifiedEmail         = "unverified-email"
	errAuthenticated           = "already-authenticated"
	// Persistent Auth Failed
	errAuth = "persistent-auth-failed"
)

// Auth User authentication handler
func Auth(c *fiber.Ctx) error {
	request := new(AuthRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}

	request.trimSpace()
	if request.AuthKey == "" {
		return new(AuthResponse).error(c, errInvalidUsernamePassword)
	}

	userID, err := auth.CanLogin(request.UserName)
	if err != nil {
		switch err {
		case auth.ErrInvalidUsername:
			return new(AuthResponse).error(c, errInvalidUsernamePassword)
		case auth.ErrUnverifiedEmail:
			return new(AuthResponse).error(c, errUnverifiedEmail)
		}
		logger.Error(err)
		return apierr.ErrSomethingWentWrong(c)
	}

	// TODO: Multiple accounts

	// Get user's login difficulty
	diff, err := auth.LoginDiff(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	// TODO: add timeout if HCaptcha disabled
	if g.Config.HCaptcha.Enabled {
		var siteKey string
		// check login difficulty
		switch diff {
		case auth.None:
		case auth.Easy:
			siteKey = g.Data.Config.HCaptcha.Easy
		case auth.Medium:
			siteKey = g.Data.Config.HCaptcha.Moderate
		case auth.Hard:
			siteKey = g.Data.Config.HCaptcha.Hard
		default:
			siteKey = g.Data.Config.HCaptcha.All
		}

		if siteKey != "" {
			if err := hcaptcha.Verify(hcaptcha.Values{
				Secret:  g.Data.Config.HCaptcha.Secret,
				Token:   request.Captcha,
				SiteKey: siteKey,
				IP:      clientip.Get(c),
			}); err != nil {
				if err != hcaptcha.ErrInvalidToken {
					logger.Error(err)
					return apierr.ErrSomethingWentWrong(c)
				}
				return new(AuthResponse).error(c, errInvalidCaptcha, siteKey)
			}
		}
	}

	privateArmored, err := auth.AuthKeyLogin(c, userID, request.AuthKey)
	// Login
	if err != nil {
		if err != auth.ErrInvalidAuthKey {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return new(AuthResponse).error(c, errInvalidUsernamePassword)
	}

	// Get db username
	currentUser, err := user.GetPublic(session.Context(c), userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	if request.Remember {
		// Set persistent auth token
		if err := token.SetAuth(c, userID); logger.Error(err) {
			// successful failure
			return fiberutil.JSON(c, 200, AuthResponse{Success: true, PrivateArmored: privateArmored, User: currentUser, Error: errAuth})
		}
	}

	return fiberutil.JSON(c, 200, AuthResponse{Success: true, PrivateArmored: privateArmored, User: currentUser})
}
