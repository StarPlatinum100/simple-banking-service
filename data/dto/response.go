package dto

import "github.com/banking-service/data/model"

type AccountDetailsResponse struct {
	AvailableBalance float64 `json:"available_balance" binding:"required"`
	ReservedBalance  float64 `json:"reserved_balance" binding:"required"`
	Locked           bool `json:"locked" binding:"required"`
	Status           model.AccountStatus `json:"status" binding:"required"`
	AccountNumber    string `json:"account_number" binding:"required"`
	Type             model.AccountType `json:"type" binding:"required"`
}
