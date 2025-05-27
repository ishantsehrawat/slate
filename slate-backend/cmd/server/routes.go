package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishant/slate-backend/internal/handlers"
	"github.com/ishant/slate-backend/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to Slate Backend!")
    })

    api := app.Group("/api")
	auth := app.Group("/auth")
    journals := api.Group("/journals", middleware.RequireAuth)

	auth.Get("/google/login", handlers.GoogleLogin)
	auth.Get("/google/callback", handlers.GoogleCallback)

    journals.Get("/", handlers.GetMyJournals)
    journals.Get("/:id", handlers.GetJournal)
    journals.Post("/", handlers.CreateJournal)
    journals.Put("/:id", handlers.UpdateJournal)
    journals.Delete("/:id", handlers.DeleteJournal)
}
