// Code generated by qtc from "actionCSS.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line web/template/email/partials/actionCSS.qtpl:2
package partials

//line web/template/email/partials/actionCSS.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line web/template/email/partials/actionCSS.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line web/template/email/partials/actionCSS.qtpl:2
func StreamActionCSS(qw422016 *qt422016.Writer) {
//line web/template/email/partials/actionCSS.qtpl:2
	qw422016.N().S(`<style>html {background: #1a1c1e;}body {font: 600 17px \'Roboto\', Helvetica, Arial, sans-serif;color: #e6e6e6;background: #121212;padding-top: 16px;width: fit-content;margin: 16px auto;border-radius: 10px;}`)
//line web/template/email/partials/actionCSS.qtpl:18
	qw422016.N().S(`.t {font-size: 24px;text-align: center;color: #f2f2f2}`)
//line web/template/email/partials/actionCSS.qtpl:25
	qw422016.N().S(`.b {padding: unset;text-align: center;max-width: 400px;padding: 16px;margin: 0 auto}`)
//line web/template/email/partials/actionCSS.qtpl:34
	qw422016.N().S(`.l {margin: 10px 0}.l a {color: #fff;padding: 10px 16px;font-weight: 500;background-color: #0d629c;border-radius: 5px;cursor: pointer;text-decoration: none}.l a:hover {background-color: #0b5484}`)
//line web/template/email/partials/actionCSS.qtpl:53
	qw422016.N().S(`.s a {color: #38afff;background-color: transparent;border-radius: 0;text-decoration: none;}.s a:hover {color: #094b77;}.py-2 {padding-top: 6px!important;}.pb-2, .py-2  {padding-bottom: 6px!important;}.s {font-size: 12px;}</style>`)
//line web/template/email/partials/actionCSS.qtpl:76
}

//line web/template/email/partials/actionCSS.qtpl:76
func WriteActionCSS(qq422016 qtio422016.Writer) {
//line web/template/email/partials/actionCSS.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line web/template/email/partials/actionCSS.qtpl:76
	StreamActionCSS(qw422016)
//line web/template/email/partials/actionCSS.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line web/template/email/partials/actionCSS.qtpl:76
}

//line web/template/email/partials/actionCSS.qtpl:76
func ActionCSS() string {
//line web/template/email/partials/actionCSS.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line web/template/email/partials/actionCSS.qtpl:76
	WriteActionCSS(qb422016)
//line web/template/email/partials/actionCSS.qtpl:76
	qs422016 := string(qb422016.B)
//line web/template/email/partials/actionCSS.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line web/template/email/partials/actionCSS.qtpl:76
	return qs422016
//line web/template/email/partials/actionCSS.qtpl:76
}
