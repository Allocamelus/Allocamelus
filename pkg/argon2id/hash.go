package argon2id

import (
	"github.com/allocamelus/allocamelus/pkg/random"
	"golang.org/x/crypto/argon2"
)

// DefaultCost for argon2id
var DefaultCost = Cost{
	Time:    3,
	Memory:  128 * 1024, // 128MB
	Threads: 2,
	KeyLen:  32, // 256 bits
	SaltLen: 32, // 256 bits
}

// Hash password with argon2id
func Hash(password string, c Cost) *Password {
	return HashSalt(password, random.Bytes(int64(c.SaltLen)), c)
}

// HashSalt provide salt to hash password with argon2id
func HashSalt(password string, salt []byte, c Cost) *Password {
	c.FillEmpty()
	return &Password{
		Version: argon2.Version,
		Cost:    c,
		Salt:    salt,
		Key:     argon2.IDKey([]byte(password), salt, c.Time, c.Memory, c.Threads, c.KeyLen),
	}
}
