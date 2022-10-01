package middleware

import (
	"net/http"
	"strings"

	"github.com/arunap2509/ecommerce/enums"
	"github.com/arunap2509/ecommerce/token"
	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("token")

	if clientToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H {"error": "no token found in header"})
		ctx.Abort()
		return
	}

	claims, err := token.ValidateToken(strings.Split(clientToken, " ")[1])

	if err != "" {
		ctx.JSON(http.StatusBadRequest, gin.H {"error": err})
		ctx.Abort()
		return
	}

	ctx.Set(enums.USER_EMAIL, claims.Email)
	ctx.Set(enums.USER_ID, claims.UserId)
	ctx.Set(enums.IS_ADMIN, claims.IsAdmin)

	ctx.Next()
}