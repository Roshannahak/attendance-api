package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

//api/admin
func AdminRouters(route *gin.RouterGroup){
	route.GET("/", controllers.GetAdmins)
}