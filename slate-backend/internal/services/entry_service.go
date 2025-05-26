package services

import (
	"time"

	"github.com/ishant/slate-backend/internal/models"
)

var entries []models.JournalEntry
var nextID int64 = 1

func GetAllEntries() []models.JournalEntry {
	return entries
}

func CreateEntry(title, content string) models.JournalEntry {
	entry := models.JournalEntry{
		ID:        nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	nextID++
	entries = append(entries, entry)
	return entry
}
