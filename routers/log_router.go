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
	route.GET("/id/:studentId", controllers.GetLogsByStudentId)
	route.GET("/:logId", controllers.GetStudentLogByLogId)
}

//api/visitor/log
func VisitorLogRouters(route *gin.RouterGroup){
	route.POST("/entry", controllers.VisitorEntry)
	route.GET("/checkin", controllers.GetVisitorCheckedInList)
	route.GET("/", controllers.GetAllVisitorLogs)
	route.GET("/id/:visitorId", controllers.GetLogsByVisitorId)
	route.GET("/:logId", controllers.GetVisitorLogByLogId)
}
