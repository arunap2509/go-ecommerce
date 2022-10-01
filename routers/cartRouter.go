package routers

import (

	"github.com/arunap2509/ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func CartRouter(route *gin.Engine) {

	route.POST("/cart/add", controllers.AddToCart)
	route.POST("/cart/remove", controllers.RemoveFromCart)
	route.GET("/cart", controllers.GetUserCart)
}