package handlers

import "github.com/gofiber/fiber/v2"

// Level 2: Echo
func Echo(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid JSON body",
		})
	}
	return c.JSON(body)
}
