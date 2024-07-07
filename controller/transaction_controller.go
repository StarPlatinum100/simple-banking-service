package controller

import (
	"log"
	"net/http"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/service"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (tc *TransactionController) Deposit(ctx *gin.Context) {
	var request dto.TransferRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := tc.transactionService.Deposit(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to complete deposit"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deposit successful", "balance": transaction})
}
