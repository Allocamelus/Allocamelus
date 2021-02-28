package main

import (
	"flag"

	"github.com/allocamelus/allocamelus/internal/app"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

var configPath string

func init() {
	const (
		defaultConfig = "./config.json"
		configUsage   = "Path to config.json"
	)

	flag.StringVar(&configPath, "config", defaultConfig, configUsage)
	flag.StringVar(&configPath, "c", defaultConfig, configUsage+" (shorthand)")
	flag.Parse()
}

func main() {
	a := app.New(configPath)
	logger.Fatal(a.InitListener())
}
