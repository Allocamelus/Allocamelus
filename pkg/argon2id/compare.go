package argon2id

import (
	"crypto/subtle"
	"errors"
)

// ErrMismatchedHashAndPassword returned when no account is found
var ErrMismatchedHashAndPassword = errors.New("utils/password/compare.go: Mismatched hashedPassword and password")

// Compare Password struct with plaintext password
func (p *Password) Compare(password string) error {
	pass := HashSalt(password, p.Salt, p.Cost)
	if subtle.ConstantTimeCompare(p.Key, pass.Key) == 1 {
		return nil
	}
	return ErrMismatchedHashAndPassword
}
