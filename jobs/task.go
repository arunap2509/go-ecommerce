package jobs

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeOrderedToTransit = "order:intransit"
	TypeInTransitToDelivered = "intransit:delivered"
)

type OrderedToTransitPayload struct {
	OrderId int
}

type InTransitToDeliveredPayload struct {
	ShippingIds []int
	OrderId int
}

func NewOrderedToTransitTask(orderId int) (*asynq.Task, error) {
	payload, err := json.Marshal(OrderedToTransitPayload{OrderId: orderId})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeOrderedToTransit, payload), nil
}

func NewIntransitToDeliveredPayload(shippingIds []int, orderId int) (*asynq.Task, error) {
	payload, err := json.Marshal(InTransitToDeliveredPayload{ShippingIds: shippingIds, OrderId: orderId})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeInTransitToDelivered, payload), nil
}