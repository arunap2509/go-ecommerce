package main

import (
	"os"

	"github.com/arunap2509/ecommerce/middleware"
	"github.com/arunap2509/ecommerce/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8888"
	}
	
	routers.AuthRouter(router)
	routers.ProductRouter(router)

	router.Use(middleware.Authentication)
	routers.CartRouter(router)
	routers.UserRouter(router)
	routers.OrderRouter(router)

	router.Run()
}
