package routers

import(
	"github.com/arunap2509/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {

	route.GET("/user/info", controllers.GetUserInfo)
	route.PUT("/user", controllers.EditUser)
	route.POST("/user/address", controllers.AddAddress)
	route.PUT("/user/address", controllers.EditAddress)
	route.DELETE("/user/address/:id", controllers.DeleteAddress)

}