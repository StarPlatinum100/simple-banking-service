package service

import (
	"errors"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/data/model"
	"github.com/banking-service/repository"
	"github.com/banking-service/utils"
)

type TransactionService interface {
	Deposit(request dto.TransferRequest) (*model.Transaction, error)
	Withdraw(user model.User, request dto.TransferRequest) (*model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, accountRepo repository.AccountRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo, accountRepo: accountRepo}
}

func (ts *transactionService) Deposit(request dto.TransferRequest) (*model.Transaction, error) {
	receiverAccount, err := ts.accountRepo.FindByAccuntNumber(request.To)
	if err != nil {
		return nil, err
	}

	if receiverAccount.Status != model.AccountStatusActive {
		return nil, errors.New("receiver account not active")
	}

	if request.Amount <= 0 {
		return nil, errors.New("amount cannot be 0 or negative")
	}

	transaction := createTransaction(request.Amount, model.Credit, model.DepositExternal, request.From, request.To, receiverAccount.ID)
	transaction.PrevBalance = receiverAccount.AvailableBalance
	newAvailableBalance := receiverAccount.AvailableBalance + request.Amount
	receiverAccount.AvailableBalance = newAvailableBalance
	transaction.CurrentBalance = newAvailableBalance
	transaction.Description = request.Description

	if err := ts.transactionRepo.SaveTransaction(transaction); err != nil {
		return nil, err
	}

	if err := ts.accountRepo.UpdateAccountBalance(receiverAccount, newAvailableBalance); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (ts *transactionService) Withdraw(user model.User, request dto.TransferRequest) (*model.Transaction, error) {

	senderAccount, err := ts.accountRepo.FindByAccuntNumber(request.From)
	if err != nil {
		return nil, err
	}

	if senderAccount.UserID != user.ID {
		return nil, errors.New("action not allowed")
	}

	if request.Amount > senderAccount.AvailableBalance {
		return nil, errors.New("insufficient balance")
	}

	if senderAccount.Status != model.AccountStatusActive {
		return nil, errors.New("sender account not active")
	}

	if request.Amount <= 0 {
		return nil, errors.New("amount cannot be 0 or negative")
	}

	transaction := createTransaction(request.Amount, model.Debit, model.WithdrawalExternal, request.From, request.To, senderAccount.ID)
	transaction.PrevBalance = senderAccount.AvailableBalance
	newAvailableBalance := senderAccount.AvailableBalance - request.Amount
	senderAccount.AvailableBalance = newAvailableBalance
	transaction.CurrentBalance = newAvailableBalance
	transaction.Description = request.Description

	if err := ts.transactionRepo.SaveTransaction(transaction); err != nil {
		return nil, err
	}

	if err := ts.accountRepo.UpdateAccountBalance(senderAccount, newAvailableBalance); err != nil {
		return nil, err
	}

	return transaction, nil
}

func createTransaction(amount float64, transactionType model.TransactionType, purpose model.TransactionPurpose, senderAccount, receiverAccount string, accountID uint) *model.Transaction {
	return &model.Transaction{
		Amount:          amount,
		Type:            transactionType,
		Purpose:         purpose,
		Reference:       utils.GenerateRandomString(),
		Status:          model.Completed,
		SenderAccount:   senderAccount,
		ReceiverAccount: receiverAccount,
		AccountID:       accountID,
	}
}
