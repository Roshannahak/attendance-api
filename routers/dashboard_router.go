package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func DashboardRouters(route *gin.RouterGroup){
	route.GET("/stats", controllers.GetDashboardStats)
}