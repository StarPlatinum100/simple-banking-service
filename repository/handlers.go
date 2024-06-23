package repository

import "gorm.io/gorm"

func InitializeRepositories(db *gorm.DB) UserRepository {
	userRepository := NewUserRepository(db)
	return userRepository
}