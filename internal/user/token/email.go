package token

import (
	"strings"

	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/email"
	"github.com/allocamelus/allocamelus/web/template"
	emailTlp "github.com/allocamelus/allocamelus/web/template/email"
)

// SendEmail for token
func (t *Token) SendEmail(emailAddress string) error {
	var (
		mail     email.Email
		mailBody email.Body
		subject  string
		link     strings.Builder
		linkStr  string
	)
	link.WriteString("https://" + g.Config.Site.Domain)
	switch t.Type {
	case Email:
		subject = "Verification Email for " + g.Data.Config.Site.Name

		link.WriteString(g.Config.Path.Public.VerifyEmail + "?selector=" + t.Selector + "&token=" + t.token)
		linkStr = link.String()

		mailBody = email.Body{
			HTML: template.EmailTemplate(&emailTlp.Verify{
				SiteName: g.Config.Site.Name,
				Subject:  subject,
				Link:     linkStr,
			}),
			Plain: emailTlp.VerifyPlain(g.Config.Site.Name, linkStr),
		}
	case Reset:
		subject = "Password Reset Email for " + g.Data.Config.Site.Name

		link.WriteString(g.Config.Path.Public.ResetPassword + "?selector=" + t.Selector + "&token=" + t.token)
		linkStr = link.String()

		mailBody = email.Body{
			HTML: template.EmailTemplate(&emailTlp.ResetPassword{
				SiteName: g.Config.Site.Name,
				Subject:  subject,
				Link:     linkStr,
			}),
			Plain: emailTlp.ResetPlain(g.Config.Site.Name, linkStr),
		}
	}

	mail = email.Email{
		From:    g.Data.Config.Site.Name + " Bot <" + g.Data.Config.Mail.Sender + ">",
		To:      []string{emailAddress},
		Subject: subject,
		Body:    mailBody,
	}
	return mail.Send(g.Config.Mail)
}
