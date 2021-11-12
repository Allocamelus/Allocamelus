package config

import (
	"errors"
	"os"
	"time"

	"github.com/allocamelus/allocamelus/pkg/argon2id"
	"github.com/allocamelus/allocamelus/pkg/email"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Config configuration struct
type Config struct {
	Argon2Cost argon2id.Cost
	Cloudflare struct {
		Enabled bool
	}
	Cookie struct {
		PreFix string
	}
	Db struct {
		Net      string
		Server   string
		Name     string
		UserName string
		Password string
	}
	Dev          bool
	GoogleClient struct {
		Enabled bool
		ID      string
		Key     string
	}
	HCaptcha struct {
		Enabled  bool
		Secret   string
		Easy     string
		Moderate string
		Hard     string
		All      string
	}
	Logs struct {
		Level int8
		Dir   bool
		Path  string
	}
	Mail email.Config
	Path struct {
		PublicDir string
		MediaDir  string
		TmpDir    string
		Public    struct {
			VerifyEmail   string
			ResetPassword string
			Media         string
		}
	}
	Redis struct {
		Host     string
		User     string
		Password string
	}
	Session struct {
		MaxLife    int64
		Expiration int64
		Duration   struct {
			MaxLife    time.Duration
			Expiration time.Duration
		}
	}
	Site struct {
		Description string
		Domain      string
		Name        string
		Port        int64
		Prefork     bool
		BodyLimit   int64
		Static      bool
	}
	Ssl struct {
		Enabled bool
		Cert    string
		Key     string
		Port    int64
	}
}

// NewConfig initialize and return Config
func NewConfig(path string) *Config {
	config, err := ReadConfig(path)
	logger.Fatal(err)

	// Only check the config as master
	if !fiber.IsChild() {
		err = config.Validate()
		logger.Fatal(err)
	}

	config.fillTime()
	return config
}

func (c *Config) fillTime() {
	c.Session.Duration = struct {
		MaxLife    time.Duration
		Expiration time.Duration
	}{
		MaxLife:    secToDuration(c.Session.MaxLife),
		Expiration: secToDuration(c.Session.Expiration),
	}
}

func secToDuration(t int64) time.Duration {
	return time.Second * time.Duration(t)
}

// ReadConfig initializes Config
func ReadConfig(path string) (*Config, error) {
	configration := new(Config)
	file, err := os.Open(path)
	if err != nil {
		return &Config{}, errors.New("Error reading config @ " + path)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configration); err != nil {
		return &Config{}, err
	}
	return configration, nil
}

// ErrBadConfig invalid config
var ErrBadConfig = errors.New("Config Validation Error: Check klog out for more info")

// Validate Config & fills defaults if empty
func (c *Config) Validate() error {
	var hasErr bool

	if c.Argon2Cost.Time == 0 {
		c.Argon2Cost.Time = 3
		klog.Warning("Warning - Config: Missing/Invalid Argon2 Time Cost | Using Default (3)")
	}
	if c.Argon2Cost.Memory == 0 {
		c.Argon2Cost.Memory = 128 * 1024
		klog.Warning("Warning - Config: Missing/Invalid Argon2 Memory Cost | Using Default (128MB)")
	}
	if c.Argon2Cost.Threads == 0 {
		c.Argon2Cost.Threads = 2
		klog.Warning("Warning - Config: Missing/Invalid Argon2 Thread Cost | Using Default (2)")
	}
	if c.Argon2Cost.KeyLen == 0 {
		c.Argon2Cost.KeyLen = 32
		klog.Warning("Warning - Config: Missing/Invalid Argon2 Key Length | Using Default (32)")
	}
	if c.Argon2Cost.SaltLen == 0 {
		c.Argon2Cost.SaltLen = 32
		klog.Warning("Warning - Config: Missing/Invalid Argon2 Salt Length | Using Default (32)")
	}

	if c.Cookie.PreFix == "" {
		klog.Warning("Warning - Config: Missing Cookie Prefix")
	}

	if c.Db.Net == "" {
		klog.Error("Error - Config: Missing Database Network Protocol")
		hasErr = true
	}
	if c.Db.Server == "" {
		klog.Error("Error - Config: Missing Database Server")
		hasErr = true
	}
	if c.Db.Name == "" {
		klog.Error("Error - Config: Missing Database Name")
		hasErr = true
	}
	if c.Db.UserName == "" {
		klog.Error("Error - Config: Missing Database User Name")
		hasErr = true
	}
	if c.Db.Password == "" {
		klog.Warning("Warning - Config: Missing Database Password")
	}

	if c.Dev {
		klog.Info("Config: Dev Mode Enabled")
	}

	if c.GoogleClient.Enabled {
		if c.GoogleClient.ID == "" {
			klog.Error("Error - Config: Missing Google Client ID")
			hasErr = true
		}
		if c.GoogleClient.Key == "" {
			klog.Error("Error - Config: Missing Google Client Key")
			hasErr = true
		}
	}

	if c.HCaptcha.Enabled {
		if c.HCaptcha.Secret == "" {
			klog.Error("Error - Config: Missing HCaptcha Secret")
			hasErr = true
		}
		if c.HCaptcha.Easy == "" || c.HCaptcha.Moderate == "" || c.HCaptcha.Hard == "" || c.HCaptcha.All == "" {
			klog.Warning("Warning - Config: Missing HCaptcha Key(s)")
		}
	}

	if c.Path.MediaDir == "" {
		c.Path.MediaDir = "media/"
		klog.Warning("Warning - Config: Missing media file path | Using Default (media/)")
	}
	if c.Path.TmpDir == "" {
		c.Path.TmpDir = "/tmp/allocamelus_tmp_dir"
		klog.Warning("Warning - Config: Missing tmpDir file path | Using Default (/tmp/allocamelus_tmp_dir)")
	}
	if err := os.MkdirAll(c.Path.TmpDir, os.ModeSticky|os.ModePerm); err != nil {
		klog.Error("Error - Config: Creating "+c.Path.TmpDir+" Failed: ", err)
	}

	if c.Path.Public.VerifyEmail == "" {
		klog.Warning("Warning - Config: Missing Public Verify Email Path")
	}
	if c.Path.Public.ResetPassword == "" {
		klog.Warning("Warning - Config: Missing Public Reset Password Path")
	}
	if c.Path.Public.Media == "" {
		c.Path.Public.Media = "/media/"
		klog.Warning("Warning - Config: Missing public media file path | Using Default (/media/)")
	}

	if c.Redis.Host == "" {
		klog.Error("Error - Config: Missing Redis Host")
		hasErr = true
	}
	if c.Redis.User == "" {
		klog.Warning("Warning - Config: Missing Redis User")
	}
	if c.Redis.Password == "" {
		klog.Warning("Warning - Config: Missing Redis Password")
	}

	if c.Session.MaxLife == 0 {
		c.Session.MaxLife = 86400
		klog.Warning("Warning - Config: Missing/Invalid Session MaxLife Time | Using Default (86400)s")
	}
	if c.Session.Expiration == 0 {
		c.Session.Expiration = 900
		klog.Warning("Warning - Config: Missing/Invalid Session Expiration Time | Using Default (900)s")
	}

	if c.Site.Description == "" {
		c.Site.Description = "Site Description"
		klog.Warning("Warning - Config: Missing Site Description | Using Default (Site Description)")
	}
	if c.Site.Domain == "" {
		c.Site.Domain = "localhost"
		klog.Warning("Warning - Config: Missing/Invalid Site Domain | Using Default (localhost)")
	}
	if c.Site.BodyLimit == 0 {
		c.Site.BodyLimit = 50 * 1024 * 1024
		klog.Warning("Warning - Config: Missing/Invalid Site Body Limit | Using Default (50MB)")
	}

	if c.Mail.Enabled {
		if c.Mail.Insecure {
			klog.Warning("Warning - Config: Insecure Email Mode Enabled")
		}
		if c.Mail.Server == "" {
			klog.Error("Error - Config: Missing Mail Server")
			hasErr = true
		}
		if c.Mail.Sender == "" {
			c.Mail.Sender = "bot@" + c.Site.Domain
			klog.Warning("Warning - Config: Missing Mail Sender | Using Default (bot@Site.Domain)")
		}
		if c.Mail.Username == "" {
			klog.Error("Error - Config: Missing Mail Username")
			hasErr = true
		}
		if c.Mail.Password == "" {
			klog.Error("Error - Config: Missing Mail Password")
			hasErr = true
		}
	}

	if c.Site.Name == "" {
		c.Site.Description = "Site Name"
		klog.Warning("Warning - Config: Missing Site Name | Using Default (Site Name)")
	}
	if c.Site.Port == 0 {
		c.Site.Port = 8080
		klog.Warning("Warning - Config: Missing/Invalid Site Port | Using Default (8080)")
	}

	if c.Ssl.Enabled {
		if c.Ssl.Key == "" || c.Ssl.Cert == "" {
			klog.Error("Error - Config: Missing Ssl Key(and|or)Cert File Path")
			hasErr = true
		}
		if c.Ssl.Port == 0 {
			c.Ssl.Port = 8443
			klog.Warning("Warning - Config: Missing Ssl Port | Using Default (8443)")
		}
	}

	if hasErr {
		return ErrBadConfig
	}
	return nil
}
