package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(route *gin.RouterGroup) {
	route.DELETE("/user/:userId", controllers.RemoveUser)
	route.GET("/users", controllers.GetAllUser)
}
