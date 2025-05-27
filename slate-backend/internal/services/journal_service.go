package services

import (
	"github.com/ishant/slate-backend/internal/database"
	"github.com/ishant/slate-backend/internal/models"
)

func GetAllJournals() ([]models.Journal, error) {
    var journals []models.Journal
    err := database.DB.Order("created_at desc").Find(&journals).Error
    return journals, err
}

func GetJournalByID(journalID, userID uint) (models.Journal, error) {
    var journal models.Journal
    err := database.DB.First(&journal, "id = ? AND user_id = ?", journalID, userID).Error
    return journal, err
}

func CreateJournal(title, content string) (models.Journal, error) {
    journal := models.Journal{
        Title:   title,
        Content: content,
    }
    err := database.DB.Create(&journal).Error
    return journal, err
}

func UpdateJournal(id uint, userID uint, title, content string) (models.Journal, error) {
    var journal models.Journal
    err := database.DB.First(&journal,"id = ? AND user_ID = ?", id, userID).Error
    if err != nil {
        return journal, err
    }

    journal.Title = title
    journal.Content = content
    err = database.DB.Save(&journal).Error
    return journal, err
}

func DeleteJournal(id uint, userID uint) error {
    var journal models.Journal
    return database.DB.Delete(&journal,"id = ? AND user_id = ?", id, userID).Error
}
