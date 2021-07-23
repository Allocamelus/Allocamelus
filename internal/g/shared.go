package g

import (
	"errors"
	"regexp"
)

var (
	// ContentInvalidChars check for invalid characters
	ContentInvalidChars  = regexp.MustCompile(`^[^<>\[\]]*$`)
	ErrInvalidChars      = errors.New("invalid-characters")
	ErrInsufficientPerms = errors.New("insufficient-permissions")
)
