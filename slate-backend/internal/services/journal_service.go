package services

import (
	"errors"

	"github.com/ishant/slate-backend/internal/database"
	"github.com/ishant/slate-backend/internal/models"
)

// GetAllJournals returns all journals ordered by created_at descending
func GetAllJournals() ([]models.Journal, error) {
	var journals []models.Journal
	err := database.DB.Order("created_at desc").Find(&journals).Error
	return journals, err
}

// GetJournalByHash fetches a journal by its unique hash and user ID
func GetJournalByHash(hash string, userID uint) (models.Journal, error) {
	var journal models.Journal
	err := database.DB.First(&journal, "hash = ? AND user_id = ?", hash, userID).Error
	if err != nil {
		return journal, err
	}
	return journal, nil
}

// CreateJournal creates a new journal with title and content
// You should generate and assign the hash before calling this or inside the model hooks
func CreateJournal(userID int, title, content, hash string) (models.Journal, error) {
	journal := models.Journal{
		UserID:  userID,
		Title:   title,
		Content: content,
		Hash:    hash,
	}

	err := database.DB.Create(&journal).Error
	return journal, err
}

// UpdateJournalByHash updates a journal by its hash and userID
func UpdateJournalByHash(hash string, userID uint, title, content string) (models.Journal, error) {
	var journal models.Journal
	err := database.DB.First(&journal, "Hash = ? AND user_id = ?", hash, userID).Error
	if err != nil {
		return journal, err
	}

	journal.Title = title
	journal.Content = content

	err = database.DB.Save(&journal).Error
	return journal, err
}

// DeleteJournalByHash deletes a journal by its hash and userID
func DeleteJournalByHash(hash string, userID uint) error {
	var journal models.Journal
	// First ensure journal exists with that hash and user
	err := database.DB.First(&journal, "hash = ? AND user_id = ?", hash, userID).Error
	if err != nil {
		return errors.New("journal not found or unauthorized")
	}

	return database.DB.Delete(&journal).Error
}
