package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/arunap2509/ecommerce/enums"
	"github.com/arunap2509/ecommerce/jobs"
	"github.com/arunap2509/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

func InstantBuy(ctx *gin.Context) {
	var request models.InstantBuy
	var product models.Product
	var order models.Order
	userId := getUserId(ctx)

	tx := db.Begin()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		ctx.Abort()
		return
	}

	if err := tx.First(&product, request.ProductId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		ctx.Abort()
		return
	}

	order.UserId = userId
	order.PaymentType = request.PaymentType
	order.AddressId = request.AddressId
	order.OrderStatus = enums.ORDER_PLACED
	order.OrderDetails = append(order.OrderDetails, 
		models.OrderDetail{ProductId: product.Id, Price: product.Price, Quantity: request.Quantity})

	if err := tx.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		ctx.Abort()
		return
	}

	triggerMoveOrderedToInTransitJob(order.Id)

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "order placed successfully"})
}

func CheckOutFromCart(ctx *gin.Context) {
	var request models.CartCheckout
	var cart models.Cart
	var order models.Order
	var products []models.Product
	userId := getUserId(ctx)
	var cartProductIds []int

	tx := db.Begin()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		ctx.Abort()
		return
	}

	result := tx.Where("user_id = ?", userId).Find(&cart)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		ctx.Abort()
		return
	}

	cart.ProductIds.AssignTo(&cartProductIds)

	if result.RowsAffected == 0 || len(cartProductIds) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message" : "You havent added anything to cart yet, pls add"})
		ctx.Abort()
		return
	}

	if err := tx.Where(cartProductIds).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		ctx.Abort()
		return
	}

	order.PaymentType = request.PaymentType
	order.UserId = userId
	order.AddressId = request.AddressId
	order.OrderStatus = enums.ORDER_PLACED

	for _, product := range products {
		order.OrderDetails = append(order.OrderDetails, 
			models.OrderDetail{ProductId: product.Id, Price: product.Price, Quantity: 1})
	}

	if err := tx.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		ctx.Abort()
		return
	}

	err := triggerMoveOrderedToInTransitJob(order.Id)

	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "order placed successfully"})
}

func triggerMoveOrderedToInTransitJob(orderId int) error {
	workerClient := jobs.GetTaskClient()

	task, err := jobs.NewOrderedToTransitTask(orderId)

	if err != nil {
		return err
	}

	taskInfo, err := workerClient.Enqueue(task, asynq.MaxRetry(3), asynq.ProcessIn(30 * time.Second), asynq.Timeout(30 * time.Second))

	if err != nil {
		return err
    }

    log.Printf(" [*] Successfully enqueued task: %+v", taskInfo)

	return nil
}