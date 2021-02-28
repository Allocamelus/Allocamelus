package argon2id

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/allocamelus/allocamelus/pkg/logger"
)

// ErrInvalidHash returned when hash is not argon2id
var ErrInvalidHash = errors.New("argon2id/parse: Hash is not argon2id")

const argon2id = "argon2id"

// Parse formated password to fill Password struct
func (p *Password) Parse(password string) (string, error) {
	slice := strings.Split(password, "$")

	if slice[1] != "argon2id" || len(slice) < 5 {
		return slice[1], ErrInvalidHash
	}

	fmt.Sscanf(slice[2], "v=%d", &p.Version)
	p.Cost = Cost{}
	_, err := fmt.Sscanf(slice[3], "m=%d,t=%d,p=%d", &p.Cost.Memory, &p.Cost.Time, &p.Cost.Threads)
	if logger.Error(err) {
		return argon2id, ErrInvalidHash
	}
	p.Salt, err = base64.RawStdEncoding.DecodeString(slice[4])
	if err != nil {
		return argon2id, ErrInvalidHash
	}
	// Only if slice has key
	if len(slice) > 5 {
		p.Key, err = base64.RawStdEncoding.DecodeString(slice[5])
		if err != nil {
			return argon2id, ErrInvalidHash
		}
	} else {
		p.Key = nil
	}
	return argon2id, nil
}

// Parse formated password return password struct
func Parse(password string) (*Password, string, error) {
	p := new(Password)
	passwordType, err := p.Parse(password)
	return p, passwordType, err
}
