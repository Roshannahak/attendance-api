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
		//api/auth
		auth := api.Group("/auth")
		{
			//api/auth/admin
			adminAuth := auth.Group("/admin")
			{
				routers.AdminAuthRouters(adminAuth)
			}

			//api/auth/student
			studentAuth := auth.Group("/student")
			{
				routers.StudentAuthRouters(studentAuth)
			}

			//api/auth/visitor
			visitorAuth := auth.Group("/visitor")
			{
				routers.VisitorAuthRouters(visitorAuth)
			}
		}

		//api/visitor
		visitor := api.Group("/visitor")
		{
			routers.VisitorRouters(visitor)

			//api/visitor/log
			logs := visitor.Group("/log")
			{
				routers.VisitorLogRouters(logs)
			}
		}

		//api/student
		student := api.Group("/student")
		{
			routers.StudentRouters(student)

			//api/student/log
			logs := student.Group("/log")
			{
				routers.StudentLogRouters(logs)
			}
		}

		//api/admin
		admin := api.Group("/admin")
		{
			routers.AdminRouters(admin)
		}

		//api/room
		room := api.Group("/room")
		{
			routers.RoomRouters(room)
		}
	}

	router.Run()
}
