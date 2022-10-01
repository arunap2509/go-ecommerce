package jobs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/arunap2509/ecommerce/enums"
	"github.com/arunap2509/ecommerce/models"
	"github.com/hibiken/asynq"
)

func InvokeInTransitToDelivered(ctx context.Context, t *asynq.Task) error {
	var payload InTransitToDeliveredPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("error while parsing the payload: %v: %w", err, asynq.SkipRetry)
	}

	err := handleInTransitToDelivered(payload)

	if err != nil {
		fmt.Println("failed while moving order to delivered", err.Error())
	}

	fmt.Println("hello there, yay order is delivered")

	return nil
}

func handleInTransitToDelivered(payload InTransitToDeliveredPayload) error {
	var order models.Order

	tx := db.Begin()

	if err := tx.Table("shippings").Where("id IN ?", 
	payload.ShippingIds).Updates(map[string]interface{}{"delivery_status": int(enums.DELIVERED)}).Error; err != nil {
		fmt.Println("error while updating shipping data", err.Error())
		return err
	}

	if err := tx.Where("id = ?", payload.OrderId).Find(&order).Error; err != nil {
		fmt.Println("error while fetching order data", err.Error())
		return err
	}

	order.OrderStatus = enums.ORDER_DELIVERED

	if err := tx.Save(&order).Error; err != nil {
		fmt.Println("error while saving order data", err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}