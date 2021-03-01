package session

import "github.com/gofiber/fiber/v2"

// ToContext set user session to context
func ToContext(c *fiber.Ctx) {
	c.Locals(storeName, FromStore(c))
}

// FromContext get user session from fiber context
func FromContext(c *fiber.Ctx) *Session {
	session := c.Locals(storeName)
	if session != nil {
		return session.(*Session)
	}
	return &Session{}
}
