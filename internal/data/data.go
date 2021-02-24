package data

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/jdinabox/goutils/logger"
	"k8s.io/klog/v2"
)

// Data struct
type Data struct {
	Config   *Config
	database *sql.DB
	redis    *redis.Client
}

// NewData initialize and return Data
//
// Also inits klog
func NewData(configPath string) *Data {
	data := new(Data)

	data.Config = NewConfig(configPath)

	logger.InitKlog(data.Config.Logs.Level, data.Config.Logs.Path)

	if err := data.initDatabase(); err != nil {
		klog.Fatal("Backplate Database Error ", err)
		klog.Flush()
	}
	data.initRedis()

	return data
}
