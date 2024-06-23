package dto

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