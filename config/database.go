package config

import (
	"gsr_pos/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("gsr.db"))
	if err != nil {
		panic("Gagal Connect Database")
	}

	database.AutoMigrate(&models.Product{}, &models.Transaction{})

	DB = database
}
