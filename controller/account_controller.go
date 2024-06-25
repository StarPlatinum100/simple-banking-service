package controller

import (
	"log"
	"net/http"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/data/model"
	"github.com/banking-service/service"
	"github.com/gin-gonic/gin"
)

type AccountContoller struct {
	accountService service.AccountService
}

func NewAccountController(service service.AccountService) *AccountContoller {
	return &AccountContoller{accountService: service}
}

func (ac *AccountContoller) CreateAccount(ctx *gin.Context) {
	userData, _ := ctx.Get("user")

	user, ok := userData.(model.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	var request dto.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.accountService.CreateAccount(user.ID, request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save acoount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
}

func (ac *AccountContoller) FindAccountByAccountNumber(ctx *gin.Context) {
	accountNumber := ctx.Param("accountNumber")
	if accountNumber == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request, enter account number"})
		return
	}

	account, err := ac.accountService.GetAccountByAccountNumber(accountNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	ctx.JSON(http.StatusOK, *account)
}

func (ac *AccountContoller) UpdateAccount(ctx *gin.Context) {
	var request dto.UpdateAccountDetails
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.accountService.UpdateAccountDetails(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update acoount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account updated successfully"})
}

func (ac *AccountContoller) CloseAccount(ctx *gin.Context) {
	accountNumber := ctx.Param("accountNumber")
	if accountNumber == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request, enter account number"})
		return
	}

	if err := ac.accountService.CloseAccount(accountNumber); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete acoount"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
