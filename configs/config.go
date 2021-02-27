package configs

import (
	"errors"
	"os"

	"github.com/allocamelus/allocamelus/pkg/argon2id"
	"github.com/jdinabox/goutils/logger"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Config configuration struct
type Config struct {
	Argon2Cost argon2id.Cost
	Cookie     struct {
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
		Path  string
	}
	Mail struct {
		Enabled  bool
		Server   string
		Sender   string
		Username string
		Password string
	}
	Path struct {
		Public string
		Dist   string
	}
	Redis struct {
		Host     string
		User     string
		Password string
	}
	Session struct {
		MaxLife    int64
		Expiration int64
	}
	Site struct {
		Description string
		Domain      string
		Name        string
		Port        int64
		Prefork     bool
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

	err = config.Validate()
	logger.Fatal(err)
	return config
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
		klog.V(1).Info("Warning - Config: Missing/Invalid Argon2 Time Cost | Using Default (3)")
	}
	if c.Argon2Cost.Memory == 0 {
		c.Argon2Cost.Memory = 81920
		klog.V(1).Info("Warning - Config: Missing/Invalid Argon2 Memory Cost | Using Default (81920)")
	}
	if c.Argon2Cost.Threads == 0 {
		c.Argon2Cost.Threads = 2
		klog.V(1).Info("Warning - Config: Missing/Invalid Argon2 Thread Cost | Using Default (2)")
	}
	if c.Argon2Cost.KeyLen == 0 {
		c.Argon2Cost.KeyLen = 32
		klog.V(1).Info("Warning - Config: Missing/Invalid Argon2 Key Length | Using Default (32)")
	}
	if c.Argon2Cost.SaltLen == 0 {
		c.Argon2Cost.SaltLen = 32
		klog.V(1).Info("Warning - Config: Missing/Invalid Argon2 Salt Length | Using Default (32)")
	}

	if c.Cookie.PreFix == "" {
		klog.V(1).Info("Warning - Config: Missing Cookie Prefix")
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
		klog.V(1).Info("Warning - Config: Missing Database Password")
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
			klog.V(1).Info("Warning - Config: Missing HCaptcha Key(s)")
		}
	}

	if c.Path.Public == "" {
		klog.V(1).Info("Warning - Config: Missing Public File Path")
	}
	if c.Path.Dist == "" {
		klog.V(1).Info("Warning - Config: Missing Public/dist File Path")
	}

	if c.Redis.Host == "" {
		klog.Error("Error - Config: Missing Redis Host")
		hasErr = true
	}
	if c.Redis.User == "" {
		klog.V(1).Info("Warning - Config: Missing Redis User")
	}
	if c.Redis.Password == "" {
		klog.V(1).Info("Warning - Config: Missing Redis Password")
	}

	if c.Session.MaxLife == 0 {
		c.Session.MaxLife = 86400
		klog.V(1).Info("Warning - Config: Missing/Invalid Session MaxLife Time | Using Default (86400)s")
	}
	if c.Session.Expiration == 0 {
		c.Session.Expiration = 900
		klog.V(1).Info("Warning - Config: Missing/Invalid Session Expiration Time | Using Default (900)s")
	}

	if c.Site.Description == "" {
		c.Site.Description = "Site Description"
		klog.V(1).Info("Warning - Config: Missing Site Description | Using Default (Site Description)")
	}
	if c.Site.Domain == "" {
		c.Site.Domain = "localhost"
		klog.V(1).Info("Warning - Config: Missing/Invalid Site Domain | Using Default (localhost)")
	}

	if c.Mail.Enabled {
		if c.Mail.Server == "" {
			klog.Error("Error - Config: Missing Mail Server")
			hasErr = true
		}
		if c.Mail.Sender == "" {
			c.Mail.Sender = "bot@" + c.Site.Domain
			klog.V(1).Info("Warning - Config: Missing Mail Sender | Using Default (bot@Site.Domain)")
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
		klog.V(1).Info("Warning - Config: Missing Site Name | Using Default (Site Name)")
	}
	if c.Site.Port == 0 {
		c.Site.Port = 8080
		klog.V(1).Info("Warning - Config: Missing/Invalid Site Port | Using Default (8080)")
	}

	if c.Ssl.Enabled {
		if c.Ssl.Key == "" || c.Ssl.Cert == "" {
			klog.Error("Error - Config: Missing Ssl Key(and|or)Cert File Path")
			hasErr = true
		}
		if c.Ssl.Port == 0 {
			c.Ssl.Port = 8443
			klog.V(1).Info("Warning - Config: Missing Ssl Port | Using Default (8443)")
		}
	}

	if hasErr {
		return ErrBadConfig
	}
	return nil
}
