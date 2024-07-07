package repository

import (
	"log"

	"github.com/banking-service/data/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(transaction *model.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) SaveTransaction(transaction *model.Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		log.Printf("failed to save transaction: %v", err)
		return err
	}
	return nil
}
