package routers

import (
	"attendance_api/controllers"

	"github.com/gin-gonic/gin"
)

func RoomRouters(route *gin.Engine){
	route.GET("/rooms", controllers.GetAllRooms)
	route.POST("/room", controllers.CreateRoom)
}