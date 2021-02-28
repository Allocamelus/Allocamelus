package data

import (
	"time"

	"github.com/allocamelus/allocamelus/pkg/fiberutil/session"
	"github.com/allocamelus/allocamelus/pkg/fiberutil/session/stores"
	"github.com/allocamelus/allocamelus/pkg/random"
)

// NewSessionStore initializes the Session handler
func (d *Data) NewSessionStore() *session.Store {
	return session.New(session.Store{
		MaxLife:    time.Second * time.Duration(d.Config.Session.MaxLife),
		Expiration: time.Second * time.Duration(d.Config.Session.Expiration),
		Cookie: session.Cookie{
			Name:     d.Config.Cookie.PreFix + "sid",
			Domain:   d.Config.Site.Domain,
			Path:     "/",
			Secure:   true,
			HTTPOnly: true,
			SameSite: "Lax",
		},
		Data: stores.New(d.redis),
		Key: session.Key{
			Length:    32,
			Generator: random.StringBase64,
		},
	})
}
