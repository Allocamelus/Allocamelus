package hcaptcha_test

import (
	"testing"

	"github.com/allocamelus/allocamelus/pkg/hcaptcha"
)

func TestVerify(t *testing.T) {
	testData := hcaptcha.Values{
		Secret:  "0x0000000000000000000000000000000000000000",
		Token:   "10000000-aaaa-bbbb-cccc-000000000001",
		SiteKey: "10000000-ffff-ffff-ffff-000000000001",
	}
	if err := hcaptcha.Verify(testData); err != nil {
		t.Error("Failed Hcaptcha verification (check internet connection)")
	}
}
