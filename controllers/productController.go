package controllers

import (
	"net/http"

	"github.com/arunap2509/ecommerce/models"
	"github.com/gin-gonic/gin"
)

func AddProduct(ctx *gin.Context) {
	var productRequest models.ProductRequest
	var product models.Product

	if err := ctx.BindJSON(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	product.Name = productRequest.Name
	product.Image = productRequest.Image
	product.TotalAvailable = productRequest.TotalAvailable
	product.AverageRating = productRequest.AverageRating
	product.Price = productRequest.Price

	if err := db.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": product})
}

func SearchProduct(ctx *gin.Context) {
	query := ctx.Query("q")
	var products []models.Product

	result := db.Where("name LIKE ?", "%" + query + "%").Find(&products)

	if result.Error != nil || result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound, 
			gin.H{"message": "OOPS, its seems like we dont have product for this query"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

func EditProduct(ctx *gin.Context) {
	var productToEdit models.Product
	var product models.Product

	if err := ctx.BindJSON(&productToEdit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if err := db.First(&product, productToEdit.Id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	product.Name = productToEdit.Name
	product.AverageRating = productToEdit.AverageRating
	product.Image = productToEdit.Image
	product.Price = productToEdit.Price
	product.TotalAvailable = productToEdit.TotalAvailable

	if err := db.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while saving data"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProducts(ctx *gin.Context) {
	var products []models.Product

	result := db.Find(&products)

	if result.Error != nil || result.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No record found"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}