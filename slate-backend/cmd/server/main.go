package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ishant/slate-backend/internal/handlers"
	"github.com/ishant/slate-backend/internal/utils"
)

func main() {
    // Initialize database
    utils.InitDB()

    // Create Fiber app
    app := fiber.New()

    // Root route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to the Slate backend!")
    })

    // Journal (Slate) routes
    api := app.Group("/api")
    journal := api.Group("/journals")

    journal.Get("/", handlers.GetJournals)
    journal.Get("/:id", handlers.GetJournal)
    journal.Post("/", handlers.CreateJournal)
    journal.Put("/:id", handlers.UpdateJournal)
    journal.Delete("/:id", handlers.DeleteJournal)

    // Start server
    port := ":8080"
    app.Listen(port)
}
