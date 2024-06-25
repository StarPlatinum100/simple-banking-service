package repository

import (
	"errors"
	"log"

	"github.com/banking-service/data/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindById(id uint) (*model.User, error)
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

func (r *userRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	if user.ID == 0 {
		log.Println("Error fetching user from DB: invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	return &user, nil
}

func (r *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
