package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ishant/slate-backend/internal/models"
)

var DB *gorm.DB

func InitDB() {
    var err error

    dbPath := "journal.db"
    DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    fmt.Println("✅ Connected to database:", dbPath)

    if err := DB.AutoMigrate(&models.Journal{}); err != nil {
        log.Fatalf("Auto migration failed: %v", err)
    }

    fmt.Println("📦 Database migration completed")
}
