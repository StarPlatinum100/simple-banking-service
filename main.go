package main

import (
	"github.com/banking-service/controller"
	"github.com/banking-service/initializers/database"
	"github.com/banking-service/initializers/env"
	"github.com/banking-service/middleware"
	"github.com/banking-service/repository"
	"github.com/banking-service/service"
	"github.com/gin-gonic/gin"
)

func main() {

	env.LoadEnvVariables()
	db := database.InitializeDB()
	database.MakeMigrations(db)

	// initialize repositories
	userRepo := repository.NewUserRepository(db)
	currencyRepo := repository.NewCurrencyRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// initialize services
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo, currencyRepo)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo)

	// initialize controllers
	userController := controller.NewUserController(userService)
	accountController := controller.NewAccountController(accountService)
	transactionController := controller.NewTransactionController(transactionService)

	router := gin.Default()
	apiRoute := router.Group("/api/v1")

	// unauthenticated
	apiRoute.POST("/users/register", userController.SignUp)
	apiRoute.POST("/users/login", userController.Login)

	// accounts
	apiRoute.POST("/accounts", middleware.RequireAuthentication(db), accountController.CreateAccount)
	apiRoute.GET("/accounts/:accountNumber", middleware.RequireAuthentication(db), accountController.FindAccountByAccountNumber)
	apiRoute.PUT("/accounts", middleware.RequireAuthentication(db), middleware.RequireAdminPrivilege, accountController.UpdateAccount)
	apiRoute.PUT("/accounts/:accountNumber", middleware.RequireAuthentication(db), middleware.RequireAdminPrivilege, accountController.CloseAccount)

	// transactions
	apiRoute.POST("transactions/deposit", transactionController.Deposit)


	apiRoute.GET("/ping", middleware.RequireAuthentication(db), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
