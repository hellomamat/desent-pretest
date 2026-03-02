package routes

import (
	"desent-pretest/handlers"
	"desent-pretest/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Level 1: Ping
	app.Get("/ping", handlers.Ping)

	// Level 2: Echo
	app.Post("/echo", handlers.Echo)

	// Level 3 & 6: Books (public)
	app.Post("/books", handlers.CreateBook)
	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)

	// Level 4: Update & Delete
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	// Level 5: Auth
	app.Post("/auth/token", handlers.Login)

	// Level 5: Protected books endpoint
	protected := app.Group("/protected")
	protected.Use(middleware.AuthGuard)
	protected.Get("/books", handlers.GetBooks)
}
