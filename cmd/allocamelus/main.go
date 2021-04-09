package main

import (
	"flag"
	"log"

	"github.com/allocamelus/allocamelus/internal/app"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"gopkg.in/gographics/imagick.v3/imagick"
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

	imagick.Initialize()
	defer func() {
		log.Println("Terminating imagick")
		imagick.Terminate()
	}()

	serverClosed := make(chan struct{})
	go a.AwaitAndClose(serverClosed)

	logger.Fatal(a.InitListener())
	<-serverClosed
}
