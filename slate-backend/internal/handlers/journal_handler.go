package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ishant/slate-backend/internal/models"
	"github.com/ishant/slate-backend/internal/services"
	"github.com/ishant/slate-backend/internal/utils"
	"gorm.io/gorm"
)

type CreateJournalInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GetJournals handles GET /api/journals
func GetMyJournals(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var journals []models.Journal
	db := c.Locals("db").(*gorm.DB) // if you’ve injected DB into context

	if err := db.Where("user_id = ?", userID).Find(&journals).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch journals",
		})
	}

	return c.JSON(journals)
}

// GetJournal handles GET /api/journals/:id
func GetJournal(c *fiber.Ctx) error {
    userID, errUser := utils.GetUserIDFromContext(c)
    if errUser != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }
    idParam := c.Params("id")
    id, errJournal := strconv.Atoi(idParam)
    if errJournal != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid journal ID"})
    }

    journal, err := services.GetJournalByID(uint(id), uint(userID))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Journal not found"})
    }
    return c.JSON(journal)
}

// CreateJournal handles POST /api/journals
func CreateJournal(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var input CreateJournalInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	journal := models.Journal{
		UserID:  userID,
		Title:   input.Title,
		Content: input.Content,
	}

	db := c.Locals("db").(*gorm.DB)
	if err := db.Create(&journal).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create journal",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(journal)
}

// UpdateJournal handles PUT /api/journals/:id
func UpdateJournal(c *fiber.Ctx) error {
    userID, err := utils.GetUserIDFromContext(c)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }
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

    journal, err := services.UpdateJournal(uint(id), uint(userID), input.Title, input.Content)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(journal)
}

// DeleteJournal handles DELETE /api/journals/:id
func DeleteJournal(c *fiber.Ctx) error {
    userID, err := utils.GetUserIDFromContext(c)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }
    idParam := c.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid journal ID"})
    }

    err = services.DeleteJournal(uint(id), uint(userID))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
