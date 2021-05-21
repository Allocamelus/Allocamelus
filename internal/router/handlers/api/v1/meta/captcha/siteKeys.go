package captcha

import (
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/pkg/fiberutil"
	"github.com/gofiber/fiber/v2"
)

type siteKeysResp struct {
	SiteKeys     siteKeys               `json:"site-keys"`
	Difficulties map[string]interface{} `json:"difficulties"`
}

type siteKeys struct {
	Easy     string `json:"easy"`
	Moderate string `json:"moderate"`
	Hard     string `json:"hard"`
	All      string `json:"all"`
}

// SiteKeys site key handler
func SiteKeys(c *fiber.Ctx) error {
	return fiberutil.JSON(c, 200, siteKeysResp{
		SiteKeys: siteKeys{
			Easy:     g.Config.HCaptcha.Easy,
			Moderate: g.Config.HCaptcha.Moderate,
			Hard:     g.Config.HCaptcha.Hard,
			All:      g.Config.HCaptcha.All,
		},
		Difficulties: map[string]interface{}{
			"user": map[string]string{
				"create":     "moderate",
				"emailToken": "moderate",
			},
		},
	})
}
