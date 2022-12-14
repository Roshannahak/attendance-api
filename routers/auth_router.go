package routers

import (
	"attendance_api/controllers/auth"

	"github.com/gin-gonic/gin"
)

//api/auth/student
func StudentAuthRouters(route *gin.RouterGroup) {
	route.POST("/register", controllers.StudentRegistration)
	route.POST("/login", controllers.StudentLogin)
	route.GET("/decode", controllers.DecryptStudentToken)
}

//api/auth/admin
func AdminAuthRouters(route *gin.RouterGroup){
	route.POST("/register", controllers.AdminRegistration)
	route.POST("/login", controllers.AdminLogin)
	route.GET("/decode", controllers.DecryptAdminToken)
}

//api/auth/visitor
func VisitorAuthRouters(route *gin.RouterGroup){
	route.POST("/login", controllers.VisitorLogin)
}
