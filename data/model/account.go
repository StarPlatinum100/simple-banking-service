package model
import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AvailableBalance float64       `gorm:"column:available_balance"`
	ReservedBalance  float64       `gorm:"column:reserved_balance"`
	Locked           bool          `gorm:"column:locked"`
	Status           AccountStatus `gorm:"column:status"`
	AccountNumber    string        `gorm:"column:account_number"`
	Type             AccountType   `gorm:"column:type"`
	UserID           uint          `gorm:"column:user_id"`
	CurrencyID       uint          `gorm:"column:currency_id"`
	Version          string        `gorm:"column:version;not null"`
}

func (Account) TableName() string {
	return "account"
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.Version == "" {
		a.Version = time.Now().String()
	}
	return nil
}
