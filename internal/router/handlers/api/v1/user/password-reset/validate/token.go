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
	errInvalidUniqueName = "invalid-unique-name"
	errInvalidToken      = "invalid-token"
	errExpiredToken      = "expired-token"
)

type tokenRequest struct {
	UniqueName string `json:"uniqueName" form:"uniqueName"`
	Selector   string `json:"selector" form:"selector"`
	Token      string `json:"token" form:"token"`
	// New password
	Password string `json:"password" form:"password"`
}

func (tr *tokenRequest) trim() {
	tr.UniqueName = strings.TrimSpace(tr.UniqueName)
	tr.Selector = strings.TrimSpace(tr.Selector)
	tr.Token = strings.TrimSpace(tr.Token)
	tr.Password = strings.TrimSpace(tr.Password)
}

type tokenResp struct {
	Success   bool   `json:"success"`
	BackupKey string `json:"backupKey,omitempty"`
	Error     string `json:"error,omitempty"`
}

// Token Reset Token handler
func Token(c *fiber.Ctx) error {
	request := new(tokenRequest)
	if err := c.BodyParser(request); err != nil {
		return apierr.ErrInvalidRequestParams(c)
	}
	request.trim()
	userID, err := user.GetIDByUniqueName(request.UniqueName)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error(err)
			return apierr.ErrSomthingWentWrong(c)
		}
		return err422(c, errInvalidUniqueName)
	}

	tkn, err := token.CheckWithID(request.Selector, request.Token, userID, token.Reset)
	if err != nil {
		if err == token.ErrExpiredToken {
			return err422(c, errExpiredToken)
		}
		return err422(c, errInvalidToken)
	}

	if err := user.ValidPassword(request.Password, request.UniqueName); err != nil {
		return err422(c, err.Error())
	}

	backupKey, err := user.UpdatePassword(userID, request.Password)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	publicKeys, err := key.GetPublicKeys(userID)
	if logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	if _, err := event.InsertNew(c, event.Reset, userID, publicKeys...); logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}
	// TODO: Send Password Change Emails
	if err := tkn.Delete(); logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}
	user.Logout(c)
	if err := token.DeleteAuthByID(userID); logger.Error(err) {
		return apierr.ErrSomthingWentWrong(c)
	}

	return fiberutil.JSON(c, 200, tokenResp{Success: true, BackupKey: backupKey})
}

func err422(c *fiber.Ctx, err string) error {
	return apierr.Err422(c, tokenResp{Error: err})
}
