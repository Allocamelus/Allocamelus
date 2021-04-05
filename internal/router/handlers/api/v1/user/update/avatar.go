package update

import (
	"fmt"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// Avatar Update handler
func Avatar(c *fiber.Ctx) error {
	// Get first file from form field "document":
	file, err := c.FormFile("document")
	logger.Error(err)
	// Save file to root directory:
	return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
}
