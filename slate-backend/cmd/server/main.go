package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ishant/slate-backend/internal/database"
	"github.com/ishant/slate-backend/internal/utils"
)

func main() {
    // Load environment and config
    utils.LoadEnv()

    // Initialize DB connection
    if err := database.Connect(); err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }
    defer database.Close()

    // Create Fiber app
    app := NewServer()

    // Start server in a goroutine so we can listen for shutdown signals
    go func() {
        if err := app.Listen(":" + utils.GetEnv("PORT", "8080")); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()
    log.Println("🚀 Server started")

    // Graceful shutdown on signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("⚠️ Shutdown signal received")

    if err := app.Shutdown(); err != nil {
        log.Fatalf("Error during shutdown: %v", err)
    }

    log.Println("👋 Server gracefully stopped")
}
