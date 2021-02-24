package errors

import "github.com/gofiber/fiber/v2"

// Error Create An json compatable response error
func Error(e string) fiber.Map {
	return fiber.Map{"error": e}
}
