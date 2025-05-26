package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Welcome to Slate API 🚀")
  })

  if err := app.Listen(":3000"); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}
