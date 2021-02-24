package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ServerStats adds server stats to header
func ServerStats(c *fiber.Ctx) error {
	// start timer
	start := time.Now()
	// next routes
	err := c.Next()
	// stop timer
	stop := time.Now()
	// Do something with response
	c.Append("Server-Stats", fmt.Sprintf("time=%v", stop.Sub(start).String()))
	// return stack error if exist
	return err
}
