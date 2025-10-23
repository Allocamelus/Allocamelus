package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/allocamelus/allocamelus/internal/app"
	"github.com/allocamelus/allocamelus/internal/version"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

var configPath string

func init() {
	if configPath = getEnvTrim("CONFIG_PATH"); configPath == "" {
		configPath = "./config.json"
	}

	v := flag.Bool("version", false, "Returns version")
	flag.Parse()
	if *v {
		fmt.Println(version.Version)
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

func getEnvTrim(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
