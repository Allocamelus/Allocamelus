//go:generate msgp

package event

import (
	"database/sql"
	"time"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/pkg/clientip"
	"github.com/allocamelus/allocamelus/internal/pkg/pgp"
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
	// Encrypted Info Key
	EncInfoKey string `msg:"encInfoKey"`
	info       Info
	infoKey    []byte
}

// Info struct for events
type Info struct {
	IP        string `msg:"ip"`
	UserAgent string `msg:"userAgent"`
}

var (
	preEvents *sql.Stmt
	preInsert *sql.Stmt
)

func initEvents(p data.Prepare) {
	preEvents = p(`SELECT COUNT(userEventId) FROM UserEvents WHERE eventType = ? AND userId = ? AND created > ?`)
	preInsert = p(`INSERT INTO UserEvents (userId, eventType, created, info, infoKey)
		VALUES (?, ?, ?, ?, ?)`)
}

// New User Event
func New(c *fiber.Ctx, t Types, userID int64, pubKey pgp.PublicKey) (*Event, error) {
	event := new(Event)
	event.UserID = userID
	event.Type = t
	event.Created = time.Now().Unix()
	event.info = newInfo(c)
	err := event.encryptInfo(pubKey)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// Insert event and encrypted info into db
func (e *Event) Insert() (err error) {
	_, err = preInsert.Exec(
		e.UserID, e.Type,
		e.Created, e.EncInfo,
		e.EncInfoKey,
	)
	return err
}

func (e *Event) encryptInfo(pubKey pgp.PublicKey) error {
	e.infoKey = random.Bytes(32)

	infoBytes, err := e.info.MarshalMsg(nil)
	if err != nil {
		return err
	}

	e.EncInfo = aesgcm.EncryptBase64(e.infoKey, infoBytes)

	// Encrypt info Key with user's public key
	e.EncInfoKey, err = pubKey.EncryptArmored(e.infoKey)
	return err
}

func newInfo(c *fiber.Ctx) (info Info) {
	info.IP = clientip.Get(c)
	info.UserAgent = c.Get("user-agent")
	return
}
