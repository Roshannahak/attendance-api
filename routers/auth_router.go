package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouters(route *gin.RouterGroup) {
	route.POST("/register", controllers.UserRegistration)
	route.POST("/login", controllers.UserLogin)
}
