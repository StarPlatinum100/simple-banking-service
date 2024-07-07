package controller

import (
	"log"
	"net/http"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/data/model"
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
	securityToken := ctx.GetHeader("token")
	if securityToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "security token cannot be empty"})
		return
	}

	var request dto.TransferRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := tc.transactionService.Deposit(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deposit successful", "transaction": transaction})
}

func (tc *TransactionController) Withdraw(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	u, ok := user.(model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	var request dto.TransferRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := tc.transactionService.Withdraw(u, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Withdrawal successful", "transaction": transaction})
}
