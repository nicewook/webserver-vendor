package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalf("gorm.Open fails: %v", err)
	}
	db.AutoMigrate(&Book{})
	DB = db
}
