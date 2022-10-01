package models

import (
	"github.com/satori/go.uuid"

	"github.com/jackc/pgtype"
)

type Cart struct {
	Id int `json:"-" gorm:"primary_key"`
	UserId uuid.UUID `json:"-"`
	User User `json:"-"`
	ProductIds pgtype.JSONB `json:"productIds" gorm:"type:jsonb;"`
}

type CartRequest struct {
	ProductId int `json:"productId"`
}