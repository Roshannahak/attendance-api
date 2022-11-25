package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

//api/room
func RoomRouters(route *gin.RouterGroup){
	route.GET("/", controllers.GetAllRooms)
	route.POST("/", controllers.CreateRoom)
	route.DELETE("/:roomId", controllers.DeleteRoom)
	route.PUT("/:roomId", controllers.UpdateRoom)
}