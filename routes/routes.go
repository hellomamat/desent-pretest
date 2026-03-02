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

	// Level 5: Auth
	app.Post("/auth/token", handlers.Login)

	// Books routes (protected with auth)
	books := app.Group("/books")
	books.Use(middleware.AuthGuard)
	books.Post("/", handlers.CreateBook)
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}
