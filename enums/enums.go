package enums

import (
	"golang.org/x/exp/slices"
)

type PaymentType int
type UserType int
type DeliveryStatus	int
type OrderStatus int

const (
	CASH_ON_DELIVERY PaymentType = 1
	ONLINE_PAYMENT PaymentType = 2
)

const (
	ADMIN UserType = 1
	USER UserType = 2
)

const (
	EN_ROUTE DeliveryStatus  = 1
	DELIVERED DeliveryStatus  = 2
)

const (
	USER_EMAIL string = "user_email"
	USER_ID string = "user_id"
	IS_ADMIN string = "is_admin"
)

const (
	ORDER_PLACED OrderStatus= 1
	ORDER_DELIVERED OrderStatus = 2
)

func (PaymentType) Values() []int {
	return []int {
		int(CASH_ON_DELIVERY),
		int(ONLINE_PAYMENT),
	}
}

func (p PaymentType) IsValid() bool {
	return slices.Contains(p.Values(), int(p))
}

func (DeliveryStatus) Values() []int {
	return []int {
		int(EN_ROUTE),
		int(DELIVERED),
	}
}

func (OrderStatus) Values() []int {
	return []int {
		int(ORDER_PLACED),
		int(ORDER_DELIVERED),
	}
}

func (d DeliveryStatus) IsValid() bool {
	return slices.Contains(d.Values(), int(d))
}

func (o OrderStatus) IsValid() bool {
	return slices.Contains(o.Values(), int(o))
}