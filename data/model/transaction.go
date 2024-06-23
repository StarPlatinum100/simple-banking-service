package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount          float64            `gorm:"type:numeric;column:amount"`
	Type            TransactionType    `gorm:"type:varchar(255);not null"`
	Purpose         TransactionPurpose `gorm:"type:varchar(255);not null"`
	Reference       string             `gorm:"type:varchar(255);unique"`
	Status          TransactionStatus  `gorm:"type:varchar(255);not null"`
	Description     string             `gorm:"type:text"`
	SenderAccount   string             `gorm:"column:sender_account"`
	ReceiverAccount string             `gorm:"column:receiver_account"`
	AccountID       uint               `gorm:"column:account_id"`
	PrevBalance     float64            `gorm:"type:numeric;column:prev_balance"`
	CurrentBalance  float64            `gorm:"type:numeric;column:current_balance"`
}

func (Transaction) TableName() string {
	return "transactions"
}
