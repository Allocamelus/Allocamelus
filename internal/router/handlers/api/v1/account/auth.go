package account

import (
	"database/sql"
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type PreAuthRequest struct {
}

// AuthRequest struct
type AuthRequest struct {
	UserName     string `json:"userName"`
	PasswordHash string `json:"passwordHash"`
	Remember     bool   `json:"remember"`
	Captcha      string `json:"captcha"`
}

func (t *AuthRequest) trimSpace() {
	t.UserName = strings.TrimSpace(t.UserName)
	t.PasswordHash = strings.TrimSpace(t.PasswordHash)
	t.Captcha = strings.TrimSpace(t.Captcha)
}

// AuthResp struct
type AuthResp struct {
	Success bool      `json:"success"`
	User    user.User `json:"user,omitempty"`
	Error   string    `json:"error,omitempty"`
	// Require Captcha
	Captcha string `json:"captcha,omitempty"`
}

const (
	withA10                    = "allocamelus"
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
	if request.UserName == "" || request.PasswordHash == "" {
		return authErr(c, errInvalidUsernamePassword)
	}

	// Check if user exists
	userID, err := user.GetIDByUserName(request.UserName)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return authErr(c, errInvalidUsernamePassword)
	}

	// TODO: Multiple accounts
	if user.LoggedIn(c) {
		s := user.ContextSession(c)
		if userID == s.UserID {
			// Allow user to re-auth if the session can't decrypt
			if s.CanDecrypt() {
				return apierr.Err403(c, AuthResp{Error: errAuthenticated})
			}
		}
	}

	// Check if user is Verified
	verified, err := user.IsVerified(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	if !verified {
		return authErr(c, errUnverifiedEmail)
	}

	// Get user's login difficulty
	diff, err := user.LoginDiff(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	// TODO: add timeout if HCaptcha disabled
	if g.Config.HCaptcha.Enabled {
		var siteKey string
		// check login difficulty
		switch diff {
		case user.None:
		case user.Easy:
			siteKey = g.Data.Config.HCaptcha.Easy
		case user.Medium:
			siteKey = g.Data.Config.HCaptcha.Moderate
		case user.Hard:
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
				return apierr.Err422(c, AuthResp{
					Error:   errInvalidCaptcha,
					Captcha: siteKey,
				})
			}
		}
	}
	// Login
	if err := user.PasswordLogin(c, userID, request.PasswordHash); err != nil {
		if err != user.ErrInvalidPassword {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return authErr(c, errInvalidUsernamePassword)
	}

	// Get db username
	currentUser, err := user.GetPublic(user.ContextSession(c), userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	if request.Remember {
		// Set persistent auth token
		if err := token.SetAuth(c, userID); logger.Error(err) {
			// successful failure
			return fiberutil.JSON(c, 200, AuthResp{Success: true, User: currentUser, Error: errAuth})
		}
	}

	return fiberutil.JSON(c, 200, AuthResp{Success: true, User: currentUser})
}

func authErr(c *fiber.Ctx, err string) error {
	return apierr.Err422(c, AuthResp{Error: err})
}
