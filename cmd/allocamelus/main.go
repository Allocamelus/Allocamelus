package main

import (
	"flag"
	"fmt"
	"os"

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
	v := flag.Bool("version", false, "Returns version")
	flag.Parse()
	if *v {
		fmt.Println(app.Version)
		os.Exit(0)
	}
}

func main() {
	a := app.New(configPath)

	serverClosed := make(chan struct{})
	go a.AwaitAndClose(serverClosed)

	// Log error if any
	logger.Error(a.InitListener())
	<-serverClosed
}
