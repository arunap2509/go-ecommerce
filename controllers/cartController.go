package controllers

import (
	"net/http"

	"github.com/arunap2509/ecommerce/enums"
	"github.com/arunap2509/ecommerce/extensions"
	"github.com/arunap2509/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func GetUserCart(ctx *gin.Context) {
	userId := getUserId(ctx)
	var cart models.Cart

	result := db.Where("user_id = ?", userId).Find(&cart)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": []models.Cart{}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cart})
}

func AddToCart(ctx *gin.Context) {
	var cartRequest models.CartRequest
	var cart models.Cart
	var exits models.Product

	if err := ctx.BindJSON(&cartRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	result := db.Where("id = ?", cartRequest.ProductId).Find(&exits)

	if result.Error != nil || result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusInternalServerError, 
			gin.H{"message": "something went wrong"})
		ctx.Abort()
		return
	}

	id, ok := ctx.Get(enums.USER_ID)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	userId := uuid.Must(uuid.FromString(id.(string)))

	isCartAvailable := db.Find(&cart, "user_id = ?", userId)

	if isCartAvailable.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	if isCartAvailable.RowsAffected == 0 {
		cart.UserId = userId
		cart.ProductIds.Set([]int{cartRequest.ProductId}) // can throw

		if err := db.Create(&cart).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
			ctx.Abort()
			return
		}
	} else {

		var ids []int
		cart.ProductIds.AssignTo(&ids)
		ids = append(ids, cartRequest.ProductId)
		cart.ProductIds.Set(ids) // can throw

		if err := db.Save(&cart).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully added to the cart"})
}

func RemoveFromCart(ctx *gin.Context) {
	var cartRequest models.CartRequest
	var cart models.Cart

	if err := ctx.BindJSON(&cartRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	id, ok := ctx.Get(enums.USER_ID)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	userId := uuid.Must(uuid.FromString(id.(string)))

	isCartAvailable := db.Find(&cart, "user_id = ?", userId)

	if isCartAvailable.Error != nil || isCartAvailable.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	var productIds []int

	cart.ProductIds.AssignTo(&productIds)
	productIds = extensions.Remove(productIds, cartRequest.ProductId)
	cart.ProductIds.Set(productIds)

	if err := db.Save(&cart).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully removed from the cart"})
}