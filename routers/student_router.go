package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

//api/student
func StudentRouters(route *gin.RouterGroup) {
	route.DELETE("/:studentId", controllers.RemoveStudent)
	route.GET("/", controllers.GetAllStudents)
	route.GET("/:quary", controllers.SearchStudent)
}
