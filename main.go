package main

import (
	"attendance_api/config"
	"attendance_api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDB()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			routers.AuthRouters(auth)
		}
		routers.UserRouters(api)

		routers.RoomRouters(api)

		routers.LogRouters(api)
	}

	router.Run("localhost:5252")
}
