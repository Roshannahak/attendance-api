package routers

import (
	"attendance_api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func StudentAuthRouters(route *gin.RouterGroup) {
	route.POST("/register", controllers.StudentRegistration)
	route.POST("/login", controllers.StudentLogin)
}

func AdminAuthRouters(route *gin.RouterGroup){
	route.POST("/register", controllers.AdminRegistration)
	route.POST("/login", controllers.AdminLogin)
}
