package jobs

import (
	"github.com/arunap2509/ecommerce/database"
	"github.com/arunap2509/ecommerce/initializers"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	initializers.InitEnv()
	db = database.Instance()
}