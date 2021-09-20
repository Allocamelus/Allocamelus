package hcaptcha

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"k8s.io/klog/v2"

	"github.com/valyala/fastjson"
)

// Values contains values for verification
// Required Secret & Token
type Values struct {
	Secret  string
	Token   string
	SiteKey string
	IP      string
}

var (
	errNilSecret  = errors.New("hcaptcha/verify: Error Empty Secret")
	errNilSiteKey = errors.New("hcaptcha/verify: Error Empty SiteKey")
	// ErrInvalidToken Invalid Token
	ErrInvalidToken = errors.New("hcaptcha/verify: Error Invalid Token")
)

// Verify verifies a hcaptcha token
func Verify(v Values) error {
	if v.Secret == "" {
		klog.Fatal(errNilSecret)
	}
	if v.Token == "" {
		return ErrInvalidToken
	}

	requestData := url.Values{
		"secret":   {v.Secret},
		"response": {v.Token},
	}

	if v.IP != "" {
		requestData.Set("remoteip", v.IP)
	}
	if v.SiteKey != "" {
		requestData.Set("sitekey", v.SiteKey)
	}

	res, err := http.PostForm("https://hcaptcha.com/siteverify", requestData)
	if err != nil {
		return err
	}

	byteBody, _ := io.ReadAll(res.Body)
	resp, err := fastjson.ParseBytes(byteBody)
	if err != nil {
		return err
	}

	if !resp.GetBool("success") {
		return ErrInvalidToken
	}

	return nil
}
