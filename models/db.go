package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDBConnection() {
	dbName := os.Getenv("DB_NAME")
	dbSession, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = dbSession
	db.AutoMigrate(&Class{})

}
