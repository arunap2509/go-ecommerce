package models

type Address struct {
	Id int `json:"id" gorm:"primary_key"`
	HouseNumber string `json:"houseNumber"`
	Area string `json:"area"`
	District string `json:"district"`
	PinCode string `json:"pinCode"`
	State string `json:"state"`
	Country string `json:"country"`
	UserId string `json:"-"`
	User User `json:"-"`
}

type AddressRequest struct {
	HouseNumber string `json:"houseNumber"`
	Area string `json:"area"`
	District string `json:"district"`
	PinCode string `json:"pinCode"`
	State string `json:"state"`
	Country string `json:"country"`
}

type EditAddress struct {
	Id int `json:"id" gorm:"primary_key"`
	HouseNumber string `json:"houseNumber"`
	Area string `json:"area"`
	District string `json:"district"`
	PinCode string `json:"pinCode"`
	State string `json:"state"`
	Country string `json:"country"`
}