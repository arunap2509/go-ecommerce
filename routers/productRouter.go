package routers

import (
	"github.com/arunap2509/ecommerce/controllers"
	"github.com/arunap2509/ecommerce/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRouter(route *gin.Engine) {

	route.GET("/product", controllers.GetProducts)
	route.POST("/product", middleware.Authentication, middleware.IsAdmin, controllers.AddProduct)
	route.PUT("/product", middleware.Authentication, middleware.IsAdmin, controllers.EditProduct)
	route.GET("/product/search", controllers.SearchProduct)
}