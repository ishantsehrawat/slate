package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ishant/slate-backend/internal/models"
)

var DB *gorm.DB

// Connect initializes the database connection.
func Connect() error {
	var err error

	// Ensure the data directory exists
	dataDir := "data"
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	dbPath := fmt.Sprintf("%s/journal.db", dataDir)
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("✅ Connected to database:", dbPath)

	if err := DB.AutoMigrate(&models.Journal{}, &models.User{}); err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	fmt.Println("📦 Database migration completed")
	return nil
}

// Close closes the database connection.
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting database instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
