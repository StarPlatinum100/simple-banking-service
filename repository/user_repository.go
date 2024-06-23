package repository

import (
	"log"

	"github.com/banking-service/data/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		log.Println("Error creating account:", err)
		return err
	}
	return nil
}
