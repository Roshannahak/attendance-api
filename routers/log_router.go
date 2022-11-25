package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func StudentLogRouters(route *gin.RouterGroup) {
	route.POST("/entry", controllers.StudentEntry)
	route.GET("/checkin", controllers.GetStudentCheckedInList)
	route.GET("/", controllers.GetAllStudentLogs)
	route.GET("/:studentId", controllers.GetLogsByStudentId)
}
