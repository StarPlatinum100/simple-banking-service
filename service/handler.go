package service

import "github.com/banking-service/repository"

func InitializeServices(userRepo repository.UserRepository) UserService {
	userService := NewUserService(userRepo)
	return userService
}