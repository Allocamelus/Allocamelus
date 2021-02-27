package argon2id

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// ToString returns a formated argon2id string
// $argon2id$v=19$$m={memory},t={time},p={threads}${salt}${key}
func (p *Password) ToString() string {
	key := base64.RawStdEncoding.EncodeToString(p.Key)
	salt := base64.RawStdEncoding.EncodeToString(p.Salt)
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Cost.Memory, p.Cost.Time, p.Cost.Threads, salt, key)
}

// ToStringNoKey returns a formated argon2id string without key
// $argon2id$v=19$$m={memory},t={time},p={threads}${salt}${key}
func (p *Password) ToStringNoKey() string {
	salt := base64.RawStdEncoding.EncodeToString(p.Salt)
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s", argon2.Version, p.Cost.Memory, p.Cost.Time, p.Cost.Threads, salt)
}
