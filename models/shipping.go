package models

import (
	"errors"

	"github.com/arunap2509/ecommerce/enums"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Shipping struct {
	Id int	`json:"id"`
	OrderDetailId int `json:"orderDetailId"`
	OrderDetail OrderDetail	`json:"-"`
	AddressId int 	`json:"addressId"`
	Address Address	`json:"-"`
	UserId uuid.UUID	`json:"userId"`
	User User	`json:"-"`
	DeliveryStatus	enums.DeliveryStatus	`json:"deliveryStatus"`
}

func (s *Shipping) BeforeCreate(db *gorm.DB) error {

	if !s.DeliveryStatus.IsValid() {
		return errors.New("delivery status is not valid")
	}

	return nil
}

func (s *Shipping) BeforeUpdate(db *gorm.DB) error {

	if !s.DeliveryStatus.IsValid() {
		return errors.New("delivery status is not valid")
	}
	
	return nil
}