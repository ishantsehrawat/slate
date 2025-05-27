package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	GoogleID  string    `gorm:"uniqueIndex;not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
