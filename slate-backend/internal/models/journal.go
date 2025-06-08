package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Journal represents a user's journal entry
type Journal struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    int       
	Hash      string    `gorm:"uniqueIndex"` // Unique hash identifier
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate is a GORM hook that runs before a new record is inserted.
// It generates a unique SHA256 hash for the journal entry if not already set.
func (j *Journal) BeforeCreate(tx *gorm.DB) (err error) {
	if j.Hash == "" {
		// Use UserID, current timestamp, and a random value to generate uniqueness
		data := fmt.Sprintf("%d-%d-%d", j.UserID, time.Now().UnixNano(), j.ID)
		hash := sha256.Sum256([]byte(data))
		j.Hash = hex.EncodeToString(hash[:])
	}
	return nil
}
