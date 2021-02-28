package session

import (
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/allocamelus/allocamelus/pkg/random"
	"github.com/gofiber/fiber/v2"
	"k8s.io/klog/v2"
)

// ErrNoSession returned by Data Get on no session found
var ErrNoSession = errors.New("Error No Session")

// Data interface of storage of session
type Data interface {
	// Returns session, error
	Get(key string) (*Session, error)
	Set(session *Session) error
	Delete(key string) error
}

// Cookie struct
type Cookie struct {
	Name     string
	Domain   string
	Path     string
	Secure   bool
	HTTPOnly bool
	SameSite string
}

// Key struct
type Key struct {
	Length    int
	Generator func(int) string
}

// Store session storage
type Store struct {
	MaxLife time.Duration
	// Expiration is Max inactive time
	Expiration time.Duration
	Cookie     Cookie
	Data       Data
	Key        Key
}

// New Store w/ value checking
//
// TODO: finish checking for all values
func New(store Store) *Store {
	if store.Key.Length < 16 {
		if store.Key.Length == 0 {
			klog.V(1).Info("Warning - Session: Missing/Invalid store.Key.Length value | Using Default (16)")
			store.Key.Length = 16
		} else {
			klog.Info("Warning - Session: store.Key.Length value < than 16")
		}
	}
	if store.Key.Generator == nil {
		klog.V(1).Info("Warning - Session: store.Key.Generator not set | Using Default (random.StringBase64)")
		store.Key.Generator = random.StringBase64
	}
	return &store
}

// Get session from store if no session is found
func (s *Store) Get(c *fiber.Ctx) *Session {
	key := c.Cookies(s.Cookie.Name, "")
	if len(key) < s.Key.Length {
		key = s.Key.Generator(s.Key.Length)
	}

	if session, ok := c.Locals("session").(*Session); ok {
		if !s.timesGood(session) {
			session.Regenerate()
		}
		return session
	}
	session, err := s.Data.Get(key)
	session.Init(s)
	if err != nil {
		if err != ErrNoSession {
			klog.Error(err)
		}
		session.Regenerate()
	} else {
		session.Status = NotUpdated
	}
	c.Locals("session", session)

	return session
}

// Set session to store
func (s *Store) Set(c *fiber.Ctx) {
	session, ok := c.Locals("session").(*Session)
	if ok {
		if s.timesGood(session) {
			if session.Status == NotUpdated {
				if time.Now().Add(s.Expiration).Sub(session.Expires) > time.Second*30 {
					// Update redis expiration every 30s on no updates
					session.Updated()
				}
			}
		} else {
			session.Regenerate()
		}

		if session.Status != NotUpdated {
			if session.Status == Deleted {
				c.Cookie(&fiber.Cookie{
					Name:     s.Cookie.Name,
					Value:    "",
					Domain:   s.Cookie.Domain,
					Expires:  time.Now(),
					Secure:   s.Cookie.Secure,
					HTTPOnly: s.Cookie.HTTPOnly,
					SameSite: s.Cookie.SameSite,
				})
			} else if len(session.Data) > 0 {
				if session.Status == Created {
					c.Cookie(&fiber.Cookie{
						Name:     s.Cookie.Name,
						Value:    session.Key,
						Domain:   s.Cookie.Domain,
						Secure:   s.Cookie.Secure,
						HTTPOnly: s.Cookie.HTTPOnly,
						SameSite: s.Cookie.SameSite,
					})
				}
				// Set status to session.Status = NotUpdated before store
				session.Status = NotUpdated

				logger.Error(s.Data.Set(session))
			}
		}
	}
}

var errBadTime = errors.New("checkTimes: Bad time")

func (s *Store) timesGood(session *Session) bool {
	if time.Since(session.Created) > s.MaxLife || !time.Now().Before(session.Expires) {
		return false
	}
	return true
}
