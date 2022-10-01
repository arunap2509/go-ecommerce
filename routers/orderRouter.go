package routers

import (
	"github.com/arunap2509/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRouter(route *gin.Engine) {

	route.POST("/orders/instant-buy", controllers.InstantBuy)
	route.POST("/orders/cart-checkout", controllers.CheckOutFromCart)
}