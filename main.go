package main

import (
	"github.com/banking-service/controller"
	"github.com/banking-service/initializers/database"
	"github.com/banking-service/initializers/env"
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
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	// apiRoute := router.Group("/api/v1")

	router.POST("/users/register", userController.SignUp)
	router.POST("/users/login", userController.Login)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
