package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

//api/visitor
func VisitorRouters(route *gin.RouterGroup){
	route.GET("/", controllers.GetAllVisitors)
	route.DELETE("/:visitorId", controllers.RemoveVisitor)
	route.GET("/:quary", controllers.SearchVisitor)
}