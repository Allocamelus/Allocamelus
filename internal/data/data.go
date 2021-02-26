package data

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/configs"
	"github.com/go-redis/redis/v8"
	"github.com/jdinabox/goutils/logger"
	"k8s.io/klog/v2"
)

// Data struct
type Data struct {
	Config   *configs.Config
	database *sql.DB
	redis    *redis.Client
}

// New initializes and returns Data struct
//
// Also inits klog
func New(configPath string) *Data {
	data := new(Data)

	data.Config = configs.NewConfig(configPath)

	logger.InitKlog(data.Config.Logs.Level, data.Config.Logs.Path)

	if err := data.initDatabase(); err != nil {
		klog.Fatal("Backplate Database Error ", err)
	}

	data.initRedis()

	return data
}
