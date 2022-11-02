package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func RoomRouters(route *gin.RouterGroup){
	route.GET("/rooms", controllers.GetAllRooms)
	route.POST("/room", controllers.CreateRoom)
	route.DELETE("/room/:roomId", controllers.DeleteRoom)
}