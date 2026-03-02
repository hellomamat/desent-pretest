package handlers

import "github.com/gofiber/fiber/v2"

// Level 2: Echo
func Echo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return c.Send(c.Body())
}
