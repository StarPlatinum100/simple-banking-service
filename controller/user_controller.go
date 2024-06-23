package controller

import (
	"log"
	"net/http"

	"github.com/banking-service/data/dto"
	"github.com/banking-service/service"
	"github.com/gin-gonic/gin"
)


type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (c *UserController) SignUp(ctx *gin.Context) {

	var userReq dto.UserSignupRequest
	if err := ctx.BindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.userService.CreateUser(userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (c *UserController) Login(ctx *gin.Context) {

	var request dto.LoginRequest

	// bind request body to the struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := c.userService.Login(request)

	ctx.JSON(http.StatusOK, gin.H{
		"token":   result,
		"message": "login successful",
	})

}
