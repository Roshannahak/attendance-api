package main

import (
	"attendance_api/config"
	"attendance_api/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	config.ConnectDB()

	routers.UserRouters(router)

	routers.RoomRouters(router)

	router.Run("localhost:5252")
}