package routers

import (
	"attendance_api/controllers/auth"

	"github.com/gin-gonic/gin"
)

//api/auth/student
func StudentAuthRouters(route *gin.RouterGroup) {
	route.POST("/register", controllers.StudentRegistration)
	route.POST("/login", controllers.StudentLogin)
}

//api/auth/admin
func AdminAuthRouters(route *gin.RouterGroup){
	route.POST("/register", controllers.AdminRegistration)
	route.POST("/login", controllers.AdminLogin)
}
