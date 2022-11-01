package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(route *gin.Engine) {

	route.POST("/user", controllers.CreateUser)
	route.DELETE("/user/:userId", controllers.RemoveUser)
	route.GET("/users", controllers.GetAllUser)
}
