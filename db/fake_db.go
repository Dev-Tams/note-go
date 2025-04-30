package db

import (
	"log"

	"github.com/dev-tams/note-go/models"
	"github.com/glebarez/sqlite" 
	"gorm.io/gorm"
	// _ "modernc.org/sqlite" 
)


var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
