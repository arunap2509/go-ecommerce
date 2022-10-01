package controllers

import (
	"net/http"

	"github.com/arunap2509/ecommerce/database"
	"github.com/arunap2509/ecommerce/enums"
	"github.com/arunap2509/ecommerce/initializers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	initializers.InitEnv()
	db = database.Instance()
}

func getUserId(ctx *gin.Context) uuid.UUID {
	id, ok := ctx.Get(enums.USER_ID)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
	}

	userId := uuid.Must(uuid.FromString(id.(string)))

	return userId
}