package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/banking-service/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RequireAuthentication(db *gorm.DB) gin.HandlerFunc {
	userRepo := repository.NewUserRepository(db)

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			log.Println("Error parsing JWT token:", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Invalid claims format")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		exp, ok := claims["exp"].(float64)

		if !ok || time.Now().Unix() > int64(exp) {
			log.Println("Expired JWT token")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, ok := claims["sub"].(float64)
		if !ok {
			log.Println("Invalid user ID in claims")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := userRepo.FindById(uint(userID))
		if err != nil {
			log.Println("Error fetching user from DB:", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", *user)
		ctx.Next()
	}
}

