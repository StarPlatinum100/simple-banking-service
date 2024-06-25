package service

import (
	"github.com/banking-service/data/dto"
	"github.com/banking-service/data/model"
	"github.com/banking-service/repository"
	"github.com/banking-service/utils"
)

type AccountService interface {
	CreateAccount(userID uint, request dto.CreateAccountRequest) error
	GetAccountByAccountNumber(acct string) (*dto.AccountDetailsResponse, error)
	UpdateAccountDetails(request dto.UpdateAccountDetails) error
	CloseAccount(acct string) error
}

type accountService struct {
	accountRepo  repository.AccountRepository
	currencyRepo repository.CurrencyRepository
}

func NewAccountService(repo repository.AccountRepository, currencyRepo repository.CurrencyRepository) AccountService {
	return &accountService{accountRepo: repo, currencyRepo: currencyRepo}
}

func (s *accountService) CreateAccount(userID uint, request dto.CreateAccountRequest) error {
	currency, err := s.currencyRepo.FindCurrencyBySymbol(request.Currency)
	if err != nil {
		return err
	}

	newAccount := model.Account{
		AvailableBalance: 0,
		ReservedBalance:  0,
		Locked:           false,
		Status:           model.AccountStatusActive,
		AccountNumber:    utils.GenerateAccountNumber(),
		Type:             request.AccountType,
		UserID:           userID,
		CurrencyID:       currency.ID,
	}

	err = s.accountRepo.CreateAccount(&newAccount)
	if err != nil {
		return err
	}

	return nil
}

func (s *accountService) GetAccountByAccountNumber(acct string) (*dto.AccountDetailsResponse, error) {
	account, err := s.accountRepo.FindByAccuntNumber(acct)
	if err != nil {
		return nil, err
	}

	accountDetailsResponse := dto.AccountDetailsResponse{
		AvailableBalance: account.AvailableBalance,
		ReservedBalance:  account.ReservedBalance,
		Locked:           account.Locked,
		Status:           account.Status,
		AccountNumber:    account.AccountNumber,
		Type:             account.Type,
	}

	return &accountDetailsResponse, nil
}

func (s *accountService) UpdateAccountDetails(request dto.UpdateAccountDetails) error {
	account, err := s.accountRepo.FindByAccuntNumber(request.AccountNumber)
	if err != nil {
		return err
	}

	if request.Type != "" {
		account.Type = request.Type
	}
	if request.Status != "" {
		account.Status = request.Status
	}

	err = s.accountRepo.UpdateAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (s *accountService) CloseAccount(acct string) error {

	if err := s.accountRepo.DeleteAccouunt(acct); err != nil {
		return err
	}

	return nil
}
