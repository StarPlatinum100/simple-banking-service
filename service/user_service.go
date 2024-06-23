package service

import (
	"log"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/data/model"
	"github.com/banking-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(request dto.UserSignupRequest) error
	Login(request dto.LoginRequest) bool
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(request dto.UserSignupRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error occured while hashing password: %v", err)
		return err
	}

	newUser := model.User{
		Password:    string(hashedPassword),
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Activated:   true,
		IsAdmin:     false,
	}

	err = s.userRepo.CreateUser(&newUser)
	if err != nil {
		log.Printf("unable to save use: %v", err)
		return err
	}

	return nil
}

func (userService) Login(request dto.LoginRequest) bool {
	return false
}
