// Code generated by qtc from "verify.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Verify email template. Implements BaseEmail methods.

//line web/template/email/verify.qtpl:2
package email

//line web/template/email/verify.qtpl:2
import "github.com/allocamelus/allocamelus/web/template"

//line web/template/email/verify.qtpl:3
import "github.com/allocamelus/allocamelus/web/template/email/partials"

//line web/template/email/verify.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line web/template/email/verify.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line web/template/email/verify.qtpl:6
type Verify struct {
	template.BaseEmail
	SiteName string
	Subject  string
	Link     string
}

//line web/template/email/verify.qtpl:15
func (e *Verify) StreamHeader(qw422016 *qt422016.Writer) {
//line web/template/email/verify.qtpl:15
	partials.StreamActionCSS(qw422016)
//line web/template/email/verify.qtpl:15
	qw422016.N().S(` `)
//line web/template/email/verify.qtpl:15
}

//line web/template/email/verify.qtpl:15
func (e *Verify) WriteHeader(qq422016 qtio422016.Writer) {
//line web/template/email/verify.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web/template/email/verify.qtpl:15
	e.StreamHeader(qw422016)
//line web/template/email/verify.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line web/template/email/verify.qtpl:15
}

//line web/template/email/verify.qtpl:15
func (e *Verify) Header() string {
//line web/template/email/verify.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line web/template/email/verify.qtpl:15
	e.WriteHeader(qb422016)
//line web/template/email/verify.qtpl:15
	qs422016 := string(qb422016.B)
//line web/template/email/verify.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line web/template/email/verify.qtpl:15
	return qs422016
//line web/template/email/verify.qtpl:15
}

//line web/template/email/verify.qtpl:16
func (e *Verify) StreamTitle(qw422016 *qt422016.Writer) {
//line web/template/email/verify.qtpl:16
	qw422016.E().S(e.Subject)
//line web/template/email/verify.qtpl:16
}

//line web/template/email/verify.qtpl:16
func (e *Verify) WriteTitle(qq422016 qtio422016.Writer) {
//line web/template/email/verify.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web/template/email/verify.qtpl:16
	e.StreamTitle(qw422016)
//line web/template/email/verify.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line web/template/email/verify.qtpl:16
}

//line web/template/email/verify.qtpl:16
func (e *Verify) Title() string {
//line web/template/email/verify.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line web/template/email/verify.qtpl:16
	e.WriteTitle(qb422016)
//line web/template/email/verify.qtpl:16
	qs422016 := string(qb422016.B)
//line web/template/email/verify.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line web/template/email/verify.qtpl:16
	return qs422016
//line web/template/email/verify.qtpl:16
}

//line web/template/email/verify.qtpl:17
func (e *Verify) StreamBody(qw422016 *qt422016.Writer) {
//line web/template/email/verify.qtpl:17
	qw422016.N().S(` <div class="t">`)
//line web/template/email/verify.qtpl:18
	qw422016.N().S(e.SiteName)
//line web/template/email/verify.qtpl:18
	qw422016.N().S(`</div> <div class="b"> <div class="pb-2"> Someone created an account using this email </div> <div class="pb-2"> Click Below to verify it was you </div> <div class="l"><a href="`)
//line web/template/email/verify.qtpl:22
	qw422016.N().S(e.Link)
//line web/template/email/verify.qtpl:22
	qw422016.N().S(`">Verify Email</a></div> <div class="s py-2">Link: <a href="`)
//line web/template/email/verify.qtpl:23
	qw422016.N().S(e.Link)
//line web/template/email/verify.qtpl:23
	qw422016.N().S(`">`)
//line web/template/email/verify.qtpl:23
	qw422016.N().S(e.Link)
//line web/template/email/verify.qtpl:23
	qw422016.N().S(`</a></div> <div> If you did not, you can ignore this email. </div> </div> `)
//line web/template/email/verify.qtpl:26
}

//line web/template/email/verify.qtpl:26
func (e *Verify) WriteBody(qq422016 qtio422016.Writer) {
//line web/template/email/verify.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web/template/email/verify.qtpl:26
	e.StreamBody(qw422016)
//line web/template/email/verify.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line web/template/email/verify.qtpl:26
}

//line web/template/email/verify.qtpl:26
func (e *Verify) Body() string {
//line web/template/email/verify.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line web/template/email/verify.qtpl:26
	e.WriteBody(qb422016)
//line web/template/email/verify.qtpl:26
	qs422016 := string(qb422016.B)
//line web/template/email/verify.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line web/template/email/verify.qtpl:26
	return qs422016
//line web/template/email/verify.qtpl:26
}

//line web/template/email/verify.qtpl:29
func StreamVerifyPlain(qw422016 *qt422016.Writer, siteName, link string) {
//line web/template/email/verify.qtpl:29
	qw422016.N().S(`
`)
//line web/template/email/verify.qtpl:30
	qw422016.N().S(siteName)
//line web/template/email/verify.qtpl:30
	qw422016.N().S(`

Someone created an account using this email
Use the link below to verify it was you

`)
//line web/template/email/verify.qtpl:35
	qw422016.N().S(link)
//line web/template/email/verify.qtpl:35
	qw422016.N().S(`

If you did not, you can ignore this email
`)
//line web/template/email/verify.qtpl:38
}

//line web/template/email/verify.qtpl:38
func WriteVerifyPlain(qq422016 qtio422016.Writer, siteName, link string) {
//line web/template/email/verify.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web/template/email/verify.qtpl:38
	StreamVerifyPlain(qw422016, siteName, link)
//line web/template/email/verify.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line web/template/email/verify.qtpl:38
}

//line web/template/email/verify.qtpl:38
func VerifyPlain(siteName, link string) string {
//line web/template/email/verify.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
//line web/template/email/verify.qtpl:38
	WriteVerifyPlain(qb422016, siteName, link)
//line web/template/email/verify.qtpl:38
	qs422016 := string(qb422016.B)
//line web/template/email/verify.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
//line web/template/email/verify.qtpl:38
	return qs422016
//line web/template/email/verify.qtpl:38
}
