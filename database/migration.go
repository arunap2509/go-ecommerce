package database

import (
	"github.com/arunap2509/ecommerce/models"
)

func CheckForMigration() {

	DB.AutoMigrate(&models.User{}, 
		&models.Address{}, 
		&models.Cart{}, 
		&models.Order{}, 
		&models.Product{},
		&models.OrderDetail{},
		&models.Shipping{})
}