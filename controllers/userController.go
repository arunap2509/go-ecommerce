package controllers

import (
	"net/http"
	"strconv"

	"github.com/arunap2509/ecommerce/models"
	"github.com/gin-gonic/gin"
)

func EditUser(ctx *gin.Context) {
	var userToUpdate models.UpdateUser
	var user models.User
	var existingUsers []models.User

	if err := ctx.BindJSON(&userToUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		ctx.Abort()
		return
	}

	userId := getUserId(ctx)

	userExists := db.Find(&user, "id = ?", userId)

	if userExists.Error != nil || userExists.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	response := db.Find(&existingUsers, "id <> ? AND (user_name = ? or email = ?)", 
	userId, userToUpdate.UserName, userToUpdate.Email)

	if response.Error != nil || response.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "username or email already exists", "error": response.Error})
		ctx.Abort()
		return
	}

	user.UserName = userToUpdate.UserName
	user.FirstName = userToUpdate.FirstName
	user.LastName = userToUpdate.LastName
	user.Email = userToUpdate.Email
	user.Phone = userToUpdate.Phone

	if err := db.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func GetUserInfo(ctx *gin.Context) {
	var userInfo models.UserInfo
	var user models.User
	var address []models.Address
	var cart models.Cart
	var orders []models.Order
	userId := getUserId(ctx)

	if err := db.Find(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	if err := db.Where("user_id = ?", userId).Find(&address).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	if err := db.Where("user_id = ?", userId).Find(&cart).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	if err := db.Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		ctx.Abort()
		return
	}

	userInfo.Address = address
	userInfo.Cart = cart
	userInfo.Orders = orders
	userInfo.FirstName = user.FirstName
	userInfo.LastName = user.LastName
	userInfo.UserName = user.UserName
	userInfo.Email = user.Email
	userInfo.Phone = user.Phone
	userInfo.Id = user.Id
	userInfo.IsAdmin = user.IsAdmin

	ctx.JSON(http.StatusOK, gin.H{"data": userInfo})
} 

func AddAddress(ctx *gin.Context) {
	var request models.AddressRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		ctx.Abort()
		return
	}

	userId := getUserId(ctx)
	address := models.Address{HouseNumber: request.HouseNumber, 
		Area: request.Area, District: request.District, 
		Country: request.Country, PinCode: request.PinCode, State: request.State, UserId: userId.String(),}

	if err := db.Create(&address).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong, try again"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "address saved against the user"})
}

func EditAddress(ctx *gin.Context) {
	var request models.EditAddress
	var address models.Address

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		ctx.Abort()
		return
	}

	if err := db.First(&address, request.Id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "address not found"})
		ctx.Abort()
		return
	}

	address.HouseNumber = request.HouseNumber
	address.Area = request.Area
	address.District = request.District
	address.State = request.State
	address.PinCode = request.PinCode
	address.Country = request.PinCode

	if err := db.Save(&address).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong, try again"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "address updated successfully"})
}

func DeleteAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	addressId, _ := strconv.Atoi(id)

	if err := db.Delete(&models.Address{}, addressId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong, try again"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "address deleted successfully"})
}