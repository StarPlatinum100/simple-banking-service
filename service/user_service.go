package service

import (
	"log"
	"os"
	"time"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/data/model"
	"github.com/banking-service/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(request dto.UserSignupRequest) error
	Login(request dto.LoginRequest) (*string, error)
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

func (s *userService) Login(request dto.LoginRequest) (*string, error) {
	user, err := s.userRepo.FindUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println("Invalid credentials:", err)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"iss":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
		"isAdmin": user.IsAdmin,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		log.Println("Error creating token:", err)
		return nil, err
	}

	return &tokenString, nil
}
