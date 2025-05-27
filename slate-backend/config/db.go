package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ishant/slate-backend/internal/models"
	"github.com/ishant/slate-backend/internal/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	utils.LoadEnv()

	// Get DB path from environment or fallback
	dbPath := utils.GetEnv("DB_PATH", "data/journal.db")

	// Ensure directory exists
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		log.Fatalf("❌ Failed to create data directory: %v", err)
	}

	// Connect to SQLite
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto-migrate all required models
	if err := db.AutoMigrate(&models.User{}, &models.Journal{}); err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	fmt.Println("✅ Database connected & migrated successfully.")
	return db
}
