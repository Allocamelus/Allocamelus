package main

import (
	"flag"

	"github.com/allocamelus/camel"
	"github.com/jdinabox/goutils/logger"
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
	c := camel.New(configPath)
	logger.Fatal(c.InitListener())
}
