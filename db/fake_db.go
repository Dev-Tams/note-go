package db

import (
	"log"
	"os"

	// "github.com/dev-tams/note-go/models"
  "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "modernc.org/sqlite"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	DB = db
	// DB.AutoMigrate(&models.User{})
}
