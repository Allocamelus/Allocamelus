// TODO: reuse less code with email-token
// TODO: Add backup key password reset

package validate

import (
	"database/sql"
	"strings"

	"github.com/allocamelus/allocamelus/internal/router/handlers/api/apierr"
	"github.com/allocamelus/allocamelus/internal/user"
	"github.com/allocamelus/allocamelus/internal/user/event"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/internal/user/token"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

var (
	errInvalidUserName = "invalid-user-name"
	errInvalidToken    = "invalid-token"
	errExpiredToken    = "expired-token"
)

type tokenRequest struct {
	UserName string `json:"userName" form:"userName"`
	Selector string `json:"selector" form:"selector"`
	Token    string `json:"token" form:"token"`
	// New password
	Password string `json:"password" form:"password"`
}

func (tr *tokenRequest) trim() {
	tr.UserName = strings.TrimSpace(tr.UserName)
	tr.Selector = strings.TrimSpace(tr.Selector)
	tr.Token = strings.TrimSpace(tr.Token)
	tr.Password = strings.TrimSpace(tr.Password)
}

type tokenResp struct {
	Success   bool   `json:"success"`
	BackupKey string `json:"backupKey,omitempty"`
	Error     string `json:"error,omitempty"`
}

// Token Password Reset Token handler
func Token(c *fiber.Ctx) error {
	request := new(tokenRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.trim()

	// Get userID
	userID, err := user.GetIDByUserName(request.UserName)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomethingWentWrong(c)
		}
		return err422(c, errInvalidUserName)
	}

	// Check Password Reset Token
	tkn, err := token.CheckWithID(request.Selector, request.Token, userID, token.Reset)
	if err != nil {
		if err == token.ErrExpiredToken {
			return err422(c, errExpiredToken)
		}
		return err422(c, errInvalidToken)
	}

	// Validate new password
	if err := user.ValidPassword(request.Password, request.UserName); err != nil {
		return err422(c, err.Error())
	}

	// Reset user password
	// Generates new key pair
	backupKey, err := user.UpdatePassword(userID, request.Password)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	// Get user publicKeys from the last keyRecoveryTime
	publicKeys, err := key.GetPublicKeys(userID)
	if logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	// Insert password reset event into database
	if _, err := event.InsertNew(c, event.Reset, userID, publicKeys...); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	// TODO: Send Password Change Emails
	if err := tkn.Delete(); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}
	user.Logout(c)
	if err := token.DeleteAuthByID(userID); logger.Error(err) {
		return apierr.ErrSomethingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, tokenResp{Success: true, BackupKey: backupKey})
}

func err422(c *fiber.Ctx, err string) error {
	return apierr.Err422(c, tokenResp{Error: err})
}
