package email

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/pkg/logger"
	mail "github.com/xhit/go-simple-mail/v2"
	"k8s.io/klog/v2"
)

var (
	// ErrEmailDisabled empty from
	ErrEmailDisabled = errors.New("models/email.go: Error Email Disabled")
	// ErrNilFrom empty from
	ErrNilFrom = errors.New("models/email.go: Error Empty From")
	// ErrNilTo empty to
	ErrNilTo = errors.New("models/email.go: Error Empty To")
	// ErrNilSubject empty subject
	ErrNilSubject = errors.New("models/email.go: Error Empty Subject")
	// ErrNilBody empty body
	ErrNilBody = errors.New("models/email.go: Error Empty Body")
)

// Config for email server
type Config struct {
	// Enabled
	Enabled bool
	// Insecure skips tls cert verify
	Insecure bool
	// SMTP server
	Server string
	// SMTP username
	Username string
	// SMTP password
	Password string
	// Sender email
	Sender string
	// From email must not be == to Sender
	//
	// Format: name <email>
	From string
	// ReplyTo
	//
	// Format: name <email>
	ReplyTo string
}

// Email struct
type Email struct {
	// From overwrites config.From
	// (optional)
	//
	// Format: name <email>
	From string
	// ReplyTo overwrites config.ReplyTo
	// (optional)
	//
	// Format: name <email>
	ReplyTo string
	// To 1 or more emails
	To []string
	// Email subject
	Subject string
	// Email body
	Body Body
}

// Body plain & html
type Body struct {
	Plain string
	HTML  string
}

// SendChan channel send email
func (e *Email) SendChan(err chan error, config Config) {
	err <- e.Send(config)
}

// Send email
func (e *Email) Send(config Config) error {
	if !config.Enabled {
		return ErrEmailDisabled
	}

	if err := e.check(); err != nil {
		return err
	}
	server := mail.NewSMTPClient()
	// SMTP Server
	server.Host = config.Server
	server.Port = 587
	server.Username = config.Username
	server.Password = config.Password
	server.Encryption = mail.EncryptionTLS

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 15 * time.Second
	// Timeout for send the data and wait respond
	server.SendTimeout = 15 * time.Second

	// Variable to keep alive connection
	if len(e.To) > 1 {
		server.KeepAlive = true
	}

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	if config.Insecure {
		server.TLSConfig = &tls.Config{InsecureSkipVerify: true} // skipcq: GSC-G402
	}

	// SMTP client
	smtpClient, err := server.Connect()
	if err != nil {
		klog.Error(err)
	}

	if e.From == "" {
		e.From = config.From
	}

	if e.ReplyTo == "" {
		e.ReplyTo = config.ReplyTo
	}

	for _, to := range e.To {
		email := mail.NewMSG()
		email.SetSender(config.Sender).
			AddTo(to).
			SetSubject(e.Subject)

		if e.From != "" {
			email.SetFrom(e.From)
		}

		if e.ReplyTo != "" {
			email.SetReplyTo(e.ReplyTo)
		}

		if e.Body.Plain != "" {
			email.SetBody(mail.TextPlain, e.Body.Plain)
			if e.Body.HTML != "" {
				email.AddAlternative(mail.TextHTML, e.Body.HTML)
			}
		} else {
			email.SetBody(mail.TextHTML, e.Body.HTML)
		}
		// Call Send and pass the client
		err = email.Send(smtpClient)
		// TODO: handle error
		logger.Error(err)
	}
	return nil
}

func (e *Email) check() error {
	if len(e.To) <= 0 {
		return ErrNilTo
	}
	if e.Subject == "" {
		return ErrNilSubject
	}
	if e.Body.HTML == "" && e.Body.Plain == "" {
		return ErrNilBody
	}
	return nil
}
