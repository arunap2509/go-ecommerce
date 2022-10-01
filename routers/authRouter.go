package routers

import(
	"github.com/arunap2509/ecommerce/controllers"
	"github.com/arunap2509/ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(route *gin.Engine) {

	route.POST("/signup", controllers.SignUp)
	route.POST("/login", controllers.Login)
	route.POST("/logout", middleware.Authentication, controllers.Logout)
}