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
    auth.Post("/logout", handlers.Logout)

    journals.Get("/", handlers.GetMyJournals)
    journals.Get("/:hash", handlers.GetJournal)
    journals.Post("/", handlers.CreateJournal)
    journals.Put("/:hash", handlers.UpdateJournal)
    journals.Delete("/:hash", handlers.DeleteJournal)
}
