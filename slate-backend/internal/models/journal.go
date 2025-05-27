package models

import (
	"time"
)

// Journal represents a user's journal entry
type Journal struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    int       
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
