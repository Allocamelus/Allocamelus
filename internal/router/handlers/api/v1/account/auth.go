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

type authRequest struct {
	With  string `json:"with" form:"with"`
	Token string `json:"token" form:"token"`
}

type authA10Token struct {
	UniqueName string `json:"uniqueName"`
	Password   string `json:"password"`
	Remember   bool   `json:"remember"`
	Captcha    string `json:"captcha"`
}

func (t *authA10Token) trimSpace() {
	t.UniqueName = strings.TrimSpace(t.UniqueName)
	t.Password = strings.TrimSpace(t.Password)
	t.Captcha = strings.TrimSpace(t.Captcha)
}

type authResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
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
	request := new(authRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	if request.With == withA10 {
		var authToken authA10Token
		if err := json.Unmarshal([]byte(request.Token), &authToken); err != nil {
			return apierr.Err422(c, errInvalidAuthToken)
		}
		authToken.trimSpace()
		if len(authToken.UniqueName) == 0 || len(authToken.Password) == 0 {
			return authErr(c, errInvalidUsernamePassword)
		}

		// Check if user exists
		userID, err := user.GetIDByUniqueName(authToken.UniqueName)
		if err != nil {
			if err != sql.ErrNoRows {
				logger.Error(err)
				return apierr.ErrSomthingWentWrong(c)
			}
			return authErr(c, errInvalidUsernamePassword)
		}

		// TODO: Multiple accounts
		if user.LoggedIn(c) {
			s := user.ContextSession(c)
			if userID == s.UserID {
				// Allow user to re-auth if the session can't decrypt
				if s.CanDecrypt() {
					return apierr.Err403(c, authResp{Error: errAuthenticated})
				}
			}
		}

		// Check if user is Verified
		verified, err := user.IsVerified(userID)
		if logger.Error(err) {
			return apierr.ErrSomthingWentWrong(c)
		}
		if !verified {
			return authErr(c, errUnverifiedEmail)
		}

		// Get user's login difficulty
		diff, err := user.LoginDiff(userID)
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
						return apierr.ErrSomthingWentWrong(c)
					}
					return apierr.Err422(c, authResp{
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
				return apierr.ErrSomthingWentWrong(c)
			}
			return authErr(c, errInvalidUsernamePassword)
		}

		if authToken.Remember {
			// Set persistent auth token
			if err := token.SetAuth(c, userID); logger.Error(err) {
				// successful failure
				return fiberutil.JSON(c, 200, authResp{Success: true, Error: errAuthToken})
			}
		}

		return fiberutil.JSON(c, 200, authResp{Success: true})
	}
	return apierr.Err422(c, errInvalidWith)
}

func authErr(c *fiber.Ctx, err string) error {
	return apierr.Err422(c, authResp{Error: err})
}
