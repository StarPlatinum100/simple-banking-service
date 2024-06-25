package middleware

import (
	"net/http"

	"github.com/banking-service/data/model"
	"github.com/gin-gonic/gin"
)

func RequireAdminPrivilege(ctx *gin.Context) {

	user, exists := ctx.Get("user")

	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	u, ok := user.(model.User)

	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !u.IsAdmin {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.Next()
}
