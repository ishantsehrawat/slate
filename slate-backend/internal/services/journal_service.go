package services

import (
	"github.com/ishant/slate-backend/internal/models"
	"github.com/ishant/slate-backend/internal/utils"
)

func GetAllJournals() ([]models.Journal, error) {
    var journals []models.Journal
    err := utils.DB.Order("created_at desc").Find(&journals).Error
    return journals, err
}

func GetJournalByID(id uint) (models.Journal, error) {
    var journal models.Journal
    err := utils.DB.First(&journal, id).Error
    return journal, err
}

func CreateJournal(title, content string) (models.Journal, error) {
    journal := models.Journal{
        Title:   title,
        Content: content,
    }
    err := utils.DB.Create(&journal).Error
    return journal, err
}

func UpdateJournal(id uint, title, content string) (models.Journal, error) {
    var journal models.Journal
    err := utils.DB.First(&journal, id).Error
    if err != nil {
        return journal, err
    }

    journal.Title = title
    journal.Content = content
    err = utils.DB.Save(&journal).Error
    return journal, err
}

func DeleteJournal(id uint) error {
    return utils.DB.Delete(&models.Journal{}, id).Error
}
