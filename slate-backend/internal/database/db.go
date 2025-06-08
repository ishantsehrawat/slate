package database

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ishant/slate-backend/internal/models"
)

var DB *gorm.DB

// Connect initializes the database connection and runs migrations.
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

	// Run safe migration for Journal model (handles 'hash' column)
	if err := safeJournalMigration(DB); err != nil {
		return fmt.Errorf("journal migration failed: %w", err)
	}

	// Migrate User model as usual
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("user migration failed: %w", err)
	}

	fmt.Println("📦 Database migration completed")
	return nil
}

// safeJournalMigration tries to add 'hash' column and populate hashes for existing journals.
func safeJournalMigration(db *gorm.DB) error {
	// 1. Check if 'hash' column exists
	type Result struct {
		Name string
	}
	var result []Result
	err := db.Raw("PRAGMA table_info(journals);").Scan(&result).Error
	if err != nil {
		return err
	}

	hashExists := false
	for _, col := range result {
		if col.Name == "hash" {
			hashExists = true
			break
		}
	}

	// If hash column doesn't exist, add it as nullable first
	if !hashExists {
		fmt.Println("Adding nullable 'hash' column to journals table...")
		err = db.Exec("ALTER TABLE journals ADD COLUMN hash TEXT;").Error
		if err != nil {
			return err
		}
	}

	// 2. Populate missing hashes for existing journals
	var journals []models.Journal
	err = db.Find(&journals).Error
	if err != nil {
		return err
	}

	for _, journal := range journals {
		if journal.Hash == "" {
			// Generate hash same way as model hook
			data := fmt.Sprintf("%d-%d-%d", journal.UserID, journal.CreatedAt.UnixNano(), journal.ID)
			hashBytes := sha256.Sum256([]byte(data))
			hash := hex.EncodeToString(hashBytes[:])

			err = db.Model(&models.Journal{}).
				Where("id = ?", journal.ID).
				Update("hash", hash).Error
			if err != nil {
				return err
			}
		}
	}

	// 3. Make 'hash' column NOT NULL and unique index
	// SQLite does not support altering column constraints directly,
	// so just create a unique index on 'hash'.
	// Note: Enforcing NOT NULL might require full table rebuild, skipped here for safety.

	fmt.Println("Creating unique index on 'hash' column...")
	err = db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_journals_hash ON journals(hash);").Error
	if err != nil {
		return err
	}

	// 4. Run AutoMigrate to apply any other schema changes (it won't add 'hash' again)
	err = db.AutoMigrate(&models.Journal{})
	if err != nil {
		return err
	}

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
