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
	Withdraw(request dto.TransferRequest) (*model.Transaction, error)
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

func (ts *transactionService) Withdraw(request dto.TransferRequest) (*model.Transaction, error) {
	return nil, nil
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
