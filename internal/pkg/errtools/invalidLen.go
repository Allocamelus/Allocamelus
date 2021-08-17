package errtools

import (
	"errors"
	"regexp"
	"strconv"
)

func InvalidLen(min, max int64) error {
	return errors.New("invalid-length-min" + strconv.Itoa(int(min)) + "-max" + strconv.Itoa(int(max)))
}

var (
	// ContentInvalidChars check for invalid characters
	ContentInvalidChars  = regexp.MustCompile(`^[^<>\[\]]*$`)
	ErrInvalidChars      = errors.New("invalid-characters")
	ErrInsufficientPerms = errors.New("insufficient-permissions")
)
