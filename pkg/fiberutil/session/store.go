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

// fiber.Ctx.Local name/key
const ctxName = "fiCtxS"

// Data interface of storage of session
type Data interface {
	// Returns session, error
	Get(key string) (*Session, error)
	Set(session *Session) error
	Delete(key string) error
}

// Cookie struct
type Cookie struct {
	// Name of Cookie NOT Value
	Name     string
	Domain   string
	Path     string
	Secure   bool
	HTTPOnly bool
	SameSite string
}

// Key struct
type Key struct {
	// Length is passed to generator
	// Min/default value of 16
	Length int
	// Generator generates session id
	Generator func(int) string
}

// Store session storage
type Store struct {
	// MaxLife of session since creation
	MaxLife time.Duration
	// Expiration max time session can go unused
	Expiration time.Duration
	// Session Cookie
	Cookie Cookie
	// Data interface for interacting with session store db
	Data Data
	// Key struct (Length, Generator)
	Key Key
}

// New Store w/ value checking
//
// TODO: finish checking for all values
func New(store *Store) *Store {
	// Min key length check
	if store.Key.Length < 16 {
		// No or Zero key length set
		if store.Key.Length == 0 {
			klog.Warning("Warning - Session: Missing/Invalid store.Key.Length value | Using Default (16)")
			// Use default key length
			store.Key.Length = 16
		} else {
			// Kill app if key length is low
			klog.Fatal("Fatal Error - Session: store.Key.Length value < than 16")
		}
	}

	// Missing key generator
	if store.Key.Generator == nil {
		klog.Warning("Warning - Session: store.Key.Generator not set | Using Default (random.StringBase64)")
		// Use default key generator
		store.Key.Generator = random.StringBase64
	}
	return store
}

// Get session from store if no session is found
func (s *Store) Get(c *fiber.Ctx) *Session {
	// Get key (Value) from session cookie
	key := c.Cookies(s.Cookie.Name, "")
	// Check key length as it may be empty
	if len(key) < s.Key.Length {
		key = s.Key.Generator(s.Key.Length)
	}

	// Try getting session from fiber context
	if session, ok := c.Locals(ctxName).(*Session); ok {
		// Is session old or expired
		// Regenerate if so
		if !s.timesGood(session) {
			// New session
			session.Regenerate()
		}
		return session
	}
	// Try getting session from database
	session, err := s.Data.Get(key)
	session.Init(s)
	if err != nil {
		// Is error ErrNoSession
		// Log if not
		if err != ErrNoSession {
			klog.Error(err)
		}
		// New session
		session.Regenerate()
	} else {
		// Session has not been updated yet in this context
		session.Status = NotUpdated
	}
	// Add session to fiber context
	c.Locals(ctxName, session)

	return session
}

// Set session to store
func (s *Store) Set(c *fiber.Ctx) {
	// Try getting session from fiber context
	if session, ok := c.Locals(ctxName).(*Session); ok {
		if s.timesGood(session) {
			// Update redis expiration every 30s on no updates
			// Prevents timeouts when not changing session data but using session
			if session.Status == NotUpdated {
				if time.Now().Add(s.Expiration).Sub(session.Expires) > time.Second*30 {
					session.Updated()
				}
			}
		} else {
			// New session
			session.Regenerate()
		}

		// Check if session has been changed
		if session.Status != NotUpdated {
			if session.Status == Deleted {
				// Unset cookie value (session key)
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
					// Send cookie to client
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

				// Add/update session in database
				// Log errors
				logger.Error(s.Data.Set(session))
			}
		}
	}
}

// timesGood checks session creation and expire times
func (s *Store) timesGood(session *Session) bool {
	if time.Since(session.Created) > s.MaxLife || !time.Now().Before(session.Expires) {
		return false
	}
	return true
}
