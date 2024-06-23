package database

import (
	"log"

	"github.com/banking-service/data/model"
	"gorm.io/gorm"
)

func MakeMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Currency{},
		&model.Account{},
		&model.Transaction{},
	)
	log.Println("migrations successful")
}
