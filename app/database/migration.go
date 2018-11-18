package database

import (
	"os"

	"w3tec.ch/go-qt-api/src/app/models"
)

// AutoCreateTables creates new tables in the database.
func AutoCreateTables() {
	if !GetDB().HasTable(&models.User{}) {
		GetDB().CreateTable(&models.User{})
	}
}

// AutoMigrateTables migrates table columns using GORM
func AutoMigrateTables() {
	GetDB().AutoMigrate(&models.User{})
}

// AutoDropTables drops tables in the "dev" environment.
func AutoDropTables() {
	if os.Getenv("API_ENV") == "dev" {
		GetDB().DropTableIfExists(&models.User{}, &models.User{})
	}
}
