package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func LogRouters(route *gin.RouterGroup) {
	route.POST("/checkin", controllers.CheckedIn)
	route.GET("/checkin", controllers.GetCheckedInList)
	route.GET("/logs", controllers.GetAllLogs)
}
