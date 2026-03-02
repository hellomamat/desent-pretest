package handlers

import "github.com/gofiber/fiber/v2"

// Level 1: Ping
func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}
