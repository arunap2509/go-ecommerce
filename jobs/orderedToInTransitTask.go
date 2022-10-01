package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/arunap2509/ecommerce/enums"
	"github.com/arunap2509/ecommerce/models"
	"github.com/hibiken/asynq"
	"gorm.io/gorm/clause"
)

func InvokeOrderedToInTransit(ctx context.Context, t *asynq.Task) error {
	var payload OrderedToTransitPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("error while parsing the payload: %v: %w", err, asynq.SkipRetry)
	}

	err := handleOrderedToInTransit(payload.OrderId)

	if err != nil {
		return err
	}
	
	fmt.Println("hello there, yay order is in transit now")
	return nil
}

func handleOrderedToInTransit(orderId int) error {

	var order models.Order
	var shipping []models.Shipping

	tx := db.Begin()

	if err := tx.Preload(clause.Associations).Where("id = ? AND order_status = 1 ", orderId).Find(&order).Error; err != nil {
		fmt.Println("error while moving the order from ordered to in transit", err.Error())
		return err
	}

	for _, orderDetail := range order.OrderDetails {
		shipping = append(shipping, 
			models.Shipping{OrderDetailId: orderDetail.Id, 
				AddressId: order.AddressId, 
				UserId: order.UserId,
				DeliveryStatus: enums.EN_ROUTE,})
	}

	if err := tx.Create(&shipping).Error; err != nil {
		fmt.Println("error while saving the data", err.Error())
		return err
	}

	err := triggerMoveInTransitToDelivered(shipping, orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func triggerMoveInTransitToDelivered(shippings []models.Shipping, orderId int) error {
	workerClient := GetTaskClient()
	var shippingIds []int

	for _, shipping := range shippings {
		shippingIds = append(shippingIds, shipping.Id)
	}

	task, err := NewIntransitToDeliveredPayload(shippingIds, orderId)

	if err != nil {
		fmt.Println("error creating delivered task", err.Error())
		return err
	}

	taskInfo, err := workerClient.Enqueue(task, asynq.MaxRetry(3), asynq.ProcessIn(30 * time.Second), asynq.Timeout(30 * time.Second))

	if err != nil {
		fmt.Println("error while enqueing delivered task", err.Error())
		return err
	}

	fmt.Printf(" [*] Successfully enqueued task: %+v", taskInfo)

	return nil
}