package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ishant/slate-backend/config"
	"github.com/ishant/slate-backend/internal/middleware"
)

func NewServer() *fiber.App {
    app := fiber.New(fiber.Config{
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  30 * time.Second,
    })

    db := config.ConnectDB() // or however you're initializing GORM

    // Inject DB into context
    app.Use(middleware.InjectDB(db))
    // Middleware
    app.Use(recover.New())
    app.Use(logger.New(logger.Config{
        Format:     "[${time}] ${status} - ${method} ${path}\n",
        TimeFormat: "02-Jan-2006 15:04:05",
        TimeZone:   "Local",
    }))

    // Register routes
    SetupRoutes(app)

    return app
}
