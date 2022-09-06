package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDBConnection() {
	dbSession, err := gorm.Open(sqlite.Open("dbname"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = dbSession
	db.AutoMigrate(&Class{})

}
