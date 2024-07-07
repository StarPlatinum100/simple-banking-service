package repository

import (
	"errors"
	"log"
	"time"

	"github.com/banking-service/data/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account *model.Account) error
	FindByAccuntNumber(acct string) (*model.Account, error)
	UpdateAccount(account *model.Account) error
	DeleteAccouunt(acc string) error
	UpdateAccountBalance(account *model.Account, newBalance float64) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(account *model.Account) error {
	if err := r.db.Create(account).Error; err != nil {
		log.Println("Error creating account:", err)
		return err
	}
	return nil
}

func (r *accountRepository) FindByAccuntNumber(acct string) (*model.Account, error) {
	var account model.Account
	if err := r.db.First(&account, "account_number = ?", acct).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func (r *accountRepository) UpdateAccount(account *model.Account) error {
	if err := r.db.Save(&account).Error; err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) DeleteAccouunt(acc string) error {
	if err := r.db.Where("account_number = ?", acc).Delete(&model.Account{}).Error; err != nil {
		log.Println("Error deleting account:", err)
		return err
	}

	return nil
}

func (r *accountRepository) UpdateAccountBalance(account *model.Account, newBalance float64) error {
	oldVersion := account.Version
	account.Version = time.Now().String()

	result := r.db.Model(account).Where("id = ? AND version = ?", account.ID, oldVersion).Updates(map[string]any{
		"available_balance": newBalance,
		"version":           account.Version,
	})

	if result.RowsAffected < 1 {
		log.Printf("failed to update account balance")
		return errors.New("failed to update account balance")
	}

	if err := result.Error; err != nil {
		log.Printf("failed to update account balance: %v", err)
		return err
	}

	return nil
}
