package routes

import (
	"desent-pretest/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Level 1: Ping
	app.Get("/ping", handlers.Ping)

	// Level 2: Echo
	app.Post("/echo", handlers.Echo)

	// Level 3 & 4: Books (public for now)
	app.Post("/books", handlers.CreateBook)
	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	// Level 5: Auth
	app.Post("/auth/token", handlers.Login)
}
