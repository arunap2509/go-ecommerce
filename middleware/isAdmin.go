package middleware

import (
	"net/http"

	"github.com/arunap2509/ecommerce/enums"

	"github.com/gin-gonic/gin"
)

func IsAdmin(ctx *gin.Context) {
	userType, exists := ctx.Get(enums.IS_ADMIN)

	if !exists || !userType.(bool)  {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You dont have access"})
		ctx.Abort()
	}

	ctx.Next()
}