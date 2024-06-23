package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password    string `gorm:"not null"`
	FirstName   string `gorm:"default:null;column:first_name"`
	LastName    string `gorm:"default:null;column:last_name"`
	Email       string `gorm:"unique;not null"`
	ImageUrl    string `gorm:"default:null;column:image_url"`
	PhoneNumber string `gorm:"unique;column:phone_number"`
	Activated   bool   `gorm:"not null"`
	IsAdmin     bool   `gorm:"default:false;column:is_admin"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Validate() error {
	if u.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if u.PhoneNumber == "" {
		return fmt.Errorf("phone number cannot be empty")
	}
	
	return nil
}
