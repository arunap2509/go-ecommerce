package models

import (
	"errors"
	"time"

	"github.com/arunap2509/ecommerce/enums"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id int `json:"id" gorm:"primary_key"`
	UserId uuid.UUID `json:"userId"`
	User User
	PaymentType enums.PaymentType `json:"paymentType"` 
	OrderStatus enums.OrderStatus `json:"orderStatus"`
	OrderDetails []OrderDetail	`json:"orderDetail"`
	AddressId int 	`json:"addressId"`
	Address Address 	`json:"-"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type OrderDetail struct {
	Id int	`json:"-"`
	OrderId int `json:"-"`
	Order Order	`json:"-"`
	ProductId int	`json:"productId"`
	Product	Product	`json:"-"`
	Quantity int 	`json:"quantity"`
	Price 	float32	`json:"price"`
}

type InstantBuy struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
	PaymentType enums.PaymentType `json:"paymentType"` 
	AddressId int 	`json:"addressId"`
}

type CartCheckout struct {
	PaymentType enums.PaymentType `json:"paymentType"` 
	AddressId int 	`json:"addressId"`
}

func (o *Order) BeforeCreate(db *gorm.DB) error {
	o.CreatedAt = time.Now().Local()

	if !o.PaymentType.IsValid() {
		return errors.New("payement type is not valid")
	}

	return nil
}

func (o *Order) BeforeUpdate(db *gorm.DB) error {
	o.UpdatedAt = time.Now().Local()

	if !o.PaymentType.IsValid() {
		return errors.New("payement type is not valid")
	}
	
	return nil
}