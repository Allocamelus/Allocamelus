//go:generate msgp

package event

import (
	"context"
	_ "embed"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/user/key"
	"github.com/allocamelus/allocamelus/pkg/aesgcm"
	"github.com/allocamelus/allocamelus/pkg/random"
	"github.com/gofiber/fiber/v2"
)

// Types for user events
type Types int8

const (
	// LoginAttempt Failed Login
	LoginAttempt Types = iota
	// Login Successful Login
	Login
	// Reset Password/Key Reset
	Reset
)

// Event struct
type Event struct {
	UserID  int64 `msg:"userId"`
	Type    Types `msg:"type"`
	Created int64 `msg:"created"`
	// Encrypted Info
	EncInfo string `msg:"encInfo"`
	// Encrypted Info Keys
	EncInfoKeys map[int64]string `msg:"encInfoKey"`
	info        Info
	infoKey     []byte
}

// Info struct for events
type Info struct {
	IP        string `msg:"ip"`
	UserAgent string `msg:"userAgent"`
}

// New User Event
func New(c *fiber.Ctx, t Types, userID int64, pubKeys ...*key.Public) (*Event, error) {
	event := new(Event)
	event.UserID = userID
	event.Type = t
	event.Created = time.Now().Unix()
	event.info = newInfo(c)
	err := event.encryptInfo(pubKeys...)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// InsertNew Create and insert User Event
func InsertNew(c *fiber.Ctx, t Types, userID int64, pubKeys ...*key.Public) (*Event, error) {
	event, err := New(c, t, userID, pubKeys...)
	if err != nil {
		return nil, err
	}
	err = event.Insert()
	if err != nil {
		return nil, err
	}
	return event, nil
}

// Insert event and encrypted info into db
func (e *Event) Insert() error {
	eventID, err := g.Data.Queries.InsertUserEvent(context.Background(), db.InsertUserEventParams{Userid: e.UserID, Eventtype: int16(e.Type), Created: e.Created, Info: e.EncInfo})
	if err != nil {
		return err
	}
	for keyID, infoKey := range e.EncInfoKeys {
		if err = g.Data.Queries.InsertUserEventKey(context.Background(), db.InsertUserEventKeyParams{Usereventid: eventID, Userkeyid: keyID, Infokey: infoKey}); err != nil {
			return err
		}
	}
	return nil
}

func (e *Event) encryptInfo(pubKeys ...*key.Public) error {
	e.infoKey = random.Bytes(32)

	infoBytes, err := e.info.MarshalMsg(nil)
	if err != nil {
		return err
	}

	e.EncInfo = aesgcm.EncryptBase64(e.infoKey, infoBytes)

	if len(pubKeys) == 0 {
		return errors.New("event/events: Error No Public Keys")
	}

	e.EncInfoKeys = map[int64]string{}

	for _, key := range pubKeys {
		// Encrypt info Key with user's public key
		encInfoKey, err := key.PublicArmored.EncryptArmored(e.infoKey)
		if err != nil {
			return err
		}
		e.EncInfoKeys[key.ID] = encInfoKey
	}
	return nil
}

func newInfo(c *fiber.Ctx) (info Info) {
	info.IP = clientip.Get(c)
	info.UserAgent = c.Get("user-agent")
	return
}
