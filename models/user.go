package models

import (
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id uuid.UUID 		`json:"id" gorm:"type:uuid;default:uuid_generate_v4()" gorm:"primary_key"`
	UserName string 	`json:"userName"`
	FirstName string 	`json:"firstName"`
	LastName string 	`json:"lastName"`
	Email string 		`json:"email"`
	Phone string		`json:"phone"`
	IsAdmin bool 		`json:"isAdmin" gorm:"default:false"`
	Password string		`json:"-"`
	CreatedAt time.Time	`json:"-"`
	ModifiedAt time.Time`json:"-"`
	Token string		`json:"token"`
	RefreshToken string	`json:"refreshToken"`
	IsDeleted bool 		`json:"-"`
}

type UpdateUser struct {
	UserName string 	`json:"userName"`
	FirstName string 	`json:"firstName"`
	LastName string 	`json:"lastName"`
	Email string 		`json:"email"`
	Phone string		`json:"phone"`
}

type SignInUser struct {
	UserName string `json:"userName" validate:"min:4,max:16"`
	FirstName string `json:"firstName" validate:"min:4,max:16"`
	LastName string `json:"lastName" validate:"min:2,max:16"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	IsAdmin bool 	`json:"isAdmin"`
}

type LoginUser struct {
	UserName string `json:"userName" validate:"min:4,max:16"`
	Password string `json:"password" validate:"min:8"`
}

type UserInfo struct {
	Id uuid.UUID 		`json:"id"`
	UserName string 	`json:"userName"`
	FirstName string 	`json:"firstName"`
	LastName string 	`json:"lastName"`
	Email string 		`json:"email"`
	Phone string 		`json:"phone"`
	IsAdmin bool 		`json:"isAdmin"`
	Address []Address	`json:"address"`
	Cart Cart			`json:"cart"`
	Orders []Order		`json:"order"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.CreatedAt = time.Now().Local()
	u.ModifiedAt = time.Now().Local()
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	u.ModifiedAt = time.Now().Local()
	return nil
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (UserInfo) TableName() string {
	return "users"
}