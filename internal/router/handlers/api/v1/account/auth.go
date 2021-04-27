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
	jsoniter "github.com/json-iterator/go"
)

// AuthRequest struct
type AuthRequest struct {
	With  string `json:"with" form:"with"`
	Token string `json:"token" form:"token"`
}

// AuthA10Token struct
type AuthA10Token struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
	Captcha  string `json:"captcha"`
}

func (t *AuthA10Token) trimSpace() {
	t.UserName = strings.TrimSpace(t.UserName)
	t.Password = strings.TrimSpace(t.Password)
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

var (
	json                = jsoniter.ConfigCompatibleWithStandardLibrary
	errInvalidAuthToken = apierr.New("invalid-auth-token")
	errInvalidWith      = apierr.New("invalid-with-value")
)

const (
	withA10                    = "allocamelus"
	errInvalidCaptcha          = "invalid-captcha"
	errInvalidUsernamePassword = "invalid-username-password"
	errUnverifiedEmail         = "unverified-email"
	errAuthenticated           = "already-authenticated"
	// Persistent Auth Failed
	errAuthToken = "persistent-auth-failed"
)

// Auth User authentication handler
func Auth(c *fiber.Ctx) error {
	request := new(AuthRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	if request.With == withA10 {
		var authToken AuthA10Token
		if err := json.Unmarshal([]byte(request.Token), &authToken); err != nil {
			return apierr.Err422(c, errInvalidAuthToken)
		}
		authToken.trimSpace()
		if len(authToken.UserName) == 0 || len(authToken.Password) == 0 {
			return authErr(c, errInvalidUsernamePassword)
		}

		// Check if user exists
		userID, err := user.GetIDByUserName(authToken.UserName)
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

			if len(siteKey) != 0 {
				if err := hcaptcha.Verify(hcaptcha.Values{
					Secret:  g.Data.Config.HCaptcha.Secret,
					Token:   authToken.Captcha,
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
		if err := user.PasswordLogin(c, userID, authToken.Password); err != nil {
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

		if authToken.Remember {
			// Set persistent auth token
			if err := token.SetAuth(c, userID); logger.Error(err) {
				// successful failure
				return fiberutil.JSON(c, 200, AuthResp{Success: true, User: currentUser, Error: errAuthToken})
			}
		}

		return fiberutil.JSON(c, 200, AuthResp{Success: true, User: currentUser})
	}
	return apierr.Err422(c, errInvalidWith)
}

func authErr(c *fiber.Ctx, err string) error {
	return apierr.Err422(c, AuthResp{Error: err})
}
