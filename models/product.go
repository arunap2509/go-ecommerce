package models

type Product struct {
	Id int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Image string `json:"image" gorm:"default:image/image1"`
	Price float32 `json:"price"`
	TotalAvailable float64 `json:"totalAvailable"`
	AverageRating float32 `json:"averageRating"`
}

type ProductRequest struct {
	Name string `json:"name"`
	Image string `json:"image" gorm:"default:image/image1"`
	Price float32 `json:"price"`
	TotalAvailable float64 `json:"totalAvailable"`
	AverageRating float32 `json:"averageRating"`
}