package repository

import (
	"github.com/banking-service/data/model"
	"gorm.io/gorm"
)

type CurrencyRepository interface {
	CreateCurrency(currency *model.Currency) error
	FindCurrencyBySymbol(symbol model.CurrencySymbol) (*model.Currency, error)

}

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) CurrencyRepository {
	return &currencyRepository{db: db}
}

func (r *currencyRepository) CreateCurrency(currency *model.Currency) error {
	return r.db.Create(currency).Error
}

func (r *currencyRepository) FindCurrencyBySymbol(symbol model.CurrencySymbol) (*model.Currency, error) {
	var currency model.Currency
	if err := r.db.Where("symbol = ?", symbol).First(&currency).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}
