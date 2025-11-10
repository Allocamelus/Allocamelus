package app

import (
	"crypto/tls"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/router/middleware"
	"github.com/allocamelus/allocamelus/internal/router/routes"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

// Allocamelus struct
type Allocamelus struct {
	Fiber *fiber.App
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// New Allocamelus server
func New(configPath string) *Allocamelus {
	g.Data = data.New(configPath)
	g.Session = g.Data.NewSessionStore()
	g.Config = g.Data.Config

	app := fiber.New(fiber.Config{
		Prefork:     g.Data.Config.Site.Prefork,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   int(g.Data.Config.Site.BodyLimit),
	})

	if g.Data.Config.Dev {
		app.Use(middleware.ServerStats)
	} else {
		// Panic Recover Middleware
		app.Use(recover.New())
	}

	app.Use(middleware.Session)
	app.Use(helmet.New())

	routes.Register(app)

	vips.Startup(nil)

	return &Allocamelus{
		Fiber: app,
	}
}

// AwaitAndClose waits for Interrupt and shuts down fiber
func (c *Allocamelus) AwaitAndClose(serverClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	log.Println("Shutting down Fiber")
	logger.Error(c.Fiber.Shutdown())
	log.Println("Closing datastores")
	g.Data.Close()
	log.Println("Shutting down vips")
	vips.Shutdown()
	log.Println("Flushing klog")
	klog.Flush()

	// Done
	close(serverClosed)
}

// InitListener initializes fiber
func (c *Allocamelus) InitListener() error {
	if g.Data.Config.Ssl.Enabled {
		// Load tls certificate
		cer, err := tls.LoadX509KeyPair(g.Data.Config.Ssl.Cert, g.Data.Config.Ssl.Key)
		if err != nil {
			return err
		}

		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}

		// Create custom listener
		ln, err := tls.Listen("tcp", ":"+strconv.Itoa(int(g.Data.Config.Ssl.Port)), tlsConfig)
		if err != nil {
			return err
		}

		return c.Fiber.Listener(ln)
	}
	return c.Fiber.Listen(":" + strconv.Itoa(int(g.Data.Config.Site.Port)))
}
