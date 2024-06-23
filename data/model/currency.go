package model

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Name    string         `gorm:"column:name;unique"`
	Symbol  CurrencySymbol `gorm:"column:symbol;unique"`
	Enabled bool           `gorm:"column:enabled"`
}

func (Currency) TableName() string {
	return "currency"
}
