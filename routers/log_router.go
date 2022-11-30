package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

//api/student/log
func StudentLogRouters(route *gin.RouterGroup) {
	route.POST("/entry", controllers.StudentEntry)
	route.GET("/checkin", controllers.GetStudentCheckedInList)
	route.GET("/", controllers.GetAllStudentLogs)
	route.GET("/:studentId", controllers.GetLogsByStudentId)
}

//api/visitor/log
func VisitorLogRouters(route *gin.RouterGroup){
	route.POST("/entry", controllers.VisitorEntry)
	route.GET("/checkin", controllers.GetVisitorCheckedInList)
	route.GET("/", controllers.GetAllVisitorLogs)
	route.GET("/:visitorId", controllers.GetLogsByVisitorId)
}
