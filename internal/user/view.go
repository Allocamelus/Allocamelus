package user

import (
	"errors"
)

var ErrViewMeNot = errors.New("user/view: Error can't view user")

// CanView can userId be viewed by session user
// 	return nil if userId can be viewed
func CanView(userId int64, sUser *Session) error {
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

	following, err := Following(sUser.UserID, userId)
	if err != nil {
		return err
	}
	if !following {
		return ErrViewMeNot
	}
	return nil
}
