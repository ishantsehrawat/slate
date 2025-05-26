package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ishant/slate-backend/internal/services"
)

// GetJournals handles GET /api/journals
func GetJournals(c *fiber.Ctx) error {
    journals, err := services.GetAllJournals()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(journals)
}

// GetJournal handles GET /api/journals/:id
func GetJournal(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid journal ID"})
    }

    journal, err := services.GetJournalByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Journal not found"})
    }

    return c.JSON(journal)
}

// CreateJournal handles POST /api/journals
func CreateJournal(c *fiber.Ctx) error {
    var input struct {
        Title   string `json:"title"`
        Content string `json:"content"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    journal, err := services.CreateJournal(input.Title, input.Content)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(journal)
}

// UpdateJournal handles PUT /api/journals/:id
func UpdateJournal(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid journal ID"})
    }

    var input struct {
        Title   string `json:"title"`
        Content string `json:"content"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    journal, err := services.UpdateJournal(uint(id), input.Title, input.Content)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(journal)
}

// DeleteJournal handles DELETE /api/journals/:id
func DeleteJournal(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid journal ID"})
    }

    err = services.DeleteJournal(uint(id))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
