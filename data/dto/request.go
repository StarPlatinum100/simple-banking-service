package dto

import "github.com/banking-service/data/model"

type UserSignupRequest struct {
	Password    string `json:"password" binding:"required,min=8"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"omitempty,len=11"`
}

type LoginRequest struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type CreateAccountRequest struct {
	Currency    model.CurrencySymbol `json:"currency" binding:"required"`
	AccountType model.AccountType    `json:"accountType" binding:"required"`
}

type UpdateAccountDetails struct {
	AccountNumber string              `json:"account_number" binding:"required"`
	Type          model.AccountType   `json:"type" binding:"required"`
	Status        model.AccountStatus `json:"status" binding:"required"`
}
