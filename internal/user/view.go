package user

import (
	"errors"

	"github.com/allocamelus/allocamelus/internal/user/session"
)

var ErrViewMeNot = errors.New("user/view: Error can't view user")

// CanView can userId be viewed by session user
// 	return nil if userId can be viewed
func CanView(userId int64, sUser *session.Session) error {
	t, err := GetType(userId)
	if err != nil {
		return err
	}
	if t.Public() {
		return nil
	}

	if !sUser.LoggedIn {
		return ErrViewMeNot
	}

	follow, err := Following(sUser.UserID, userId)
	if err != nil {
		return err
	}
	if !follow.Following {
		return ErrViewMeNot
	}
	return nil
}
