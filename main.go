package main

import (
	"desent-pretest/config"
	"desent-pretest/database"
	"desent-pretest/models"
	"desent-pretest/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env (ignore error if not present, e.g. in production)
	_ = godotenv.Load()

	cfg := config.Load()

	// Connect database
	database.Connect(cfg)

	// Auto-migrate
	if err := database.DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("migrations completed")

	// Setup Fiber
	app := fiber.New(fiber.Config{
		AppName: "Desent Pretest API",
	})

	app.Use(cors.New())
	app.Use(logger.New())

	// Setup routes
	routes.Setup(app)

	// Health check root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Desent Pretest API is running",
		})
	})

	log.Printf("server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
