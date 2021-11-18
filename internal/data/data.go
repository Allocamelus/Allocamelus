package data

import (
	"database/sql"

	"github.com/allocamelus/allocamelus/internal/config"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/go-redis/redis/v8"
	"k8s.io/klog/v2"
)

// Data struct
type Data struct {
	Config   *config.Config
	database *sql.DB
	redis    *redis.Client
}

// New initializes and returns Data struct
//
// Also inits klog
func New(configPath string) *Data {
	data := new(Data)

	data.Config = config.NewConfig(configPath)

	logger.InitKlog(data.Config.Logs.Level, data.Config.Logs.Dir, data.Config.Logs.Path)

	if err := data.initDatabase(); err != nil {
		klog.Fatal("Backplate Database Error ", err)
	}

	data.initRedis()

	return data
}
