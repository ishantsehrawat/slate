package handlers

import (
	"sort"
	"strconv"
	"strings"
	"time"

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

// GroupedJournalsResponse matches the desired response structure
// GetJournals handles GET /api/journals
type GroupedJournalsResponse struct {
	Label    string       `json:"label"`
	Journals []JournalDTO `json:"journals"`
}

// JournalDTO is a simplified journal for the response
type JournalDTO struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func GetMyJournals(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "100"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	search := c.Query("search", "")

	db := c.Locals("db").(*gorm.DB)

	var journals []models.Journal
	query := db.Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Offset(offset)
	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}
	if err := query.Find(&journals).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch journals"})
	}

	grouped := groupJournalsByDate(journals)

	return c.JSON(grouped)
}

func groupJournalsByDate(journals []models.Journal) []GroupedJournalsResponse {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	groups := map[string][]JournalDTO{}

	for _, journal := range journals {
		created := journal.CreatedAt
		startOfCreated := time.Date(created.Year(), created.Month(), created.Day(), 0, 0, 0, 0, created.Location())

		var label string
		switch {
		case startOfCreated.Equal(startOfToday):
			label = "Today"
		case startOfCreated.After(startOfToday.AddDate(0,0,-1)) && startOfCreated.Before(startOfToday):
			label = "Yesterday"
		case startOfCreated.After(startOfToday.AddDate(0, 0, -7)) && startOfCreated.After(startOfToday.AddDate(0,0,-1)):
			label = "Previous 7 Days"
		case startOfCreated.After(startOfToday.AddDate(0, 0, -30)) && startOfCreated.After(startOfToday.AddDate(0,0,-7)):
			label = "Previous 7 Days"
		case created.Year() == now.Year() && created.Month() == now.AddDate(0, -1, 0).Month():
			label = "Last Month"
		default:
			label = created.Format("January 2006") // to differentiate same month of different years
		}

		entry := JournalDTO{
			ID:    journal.ID,
			Title: journal.Title,
		}

		groups[label] = append(groups[label], entry)
	}

	// Fixed order for main groups
	fixedOrder := []string{"Today", "Yesterday", "Previous 7 Days","Previous 30 Days", "Last Month"}
	seen := map[string]bool{}
	var result []GroupedJournalsResponse

	// Add fixed groups if present
	for _, label := range fixedOrder {
		if journals, ok := groups[label]; ok {
			result = append(result, GroupedJournalsResponse{
				Label:    label,
				Journals: journals,
			})
			seen[label] = true
		}
	}

	// Collect month labels and sort descending by date
	var monthLabels []string
	for label := range groups {
		if !seen[label] {
			monthLabels = append(monthLabels, label)
		}
	}

	sort.Slice(monthLabels, func(i, j int) bool {
		ti, errI := time.Parse("January 2006", monthLabels[i])
		tj, errJ := time.Parse("January 2006", monthLabels[j])
		if errI != nil || errJ != nil {
			// fallback lex order descending if parsing fails
			return monthLabels[i] > monthLabels[j]
		}
		return ti.After(tj)
	})

	for _, label := range monthLabels {
		result = append(result, GroupedJournalsResponse{
			Label:    label,
			Journals: groups[label],
		})
	}

	return result
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

	title := input.Title
	if strings.TrimSpace(title) == "" {
		title = utils.GenerateHeadingFromContent(input.Content)
	}

	journal, err := services.UpdateJournal(uint(id), uint(userID), title, input.Content)
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