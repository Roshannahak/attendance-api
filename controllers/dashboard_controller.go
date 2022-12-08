package controllers

import (
	"attendance_api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDashboardStats(c *gin.Context) {
	studentCount, _ := studentCollection.CountDocuments(context.TODO(), bson.M{})
	visitorCount, _ := visitorCollection.CountDocuments(context.TODO(), bson.M{})
	adminCount, _ := AdminCollection.CountDocuments(context.TODO(), bson.M{})
	studentLogCount, _ := studentLogCollection.CountDocuments(context.TODO(), bson.M{})
	visitorLogCount, _ := visitorLogCollection.CountDocuments(context.TODO(), bson.M{})
	roomCount, _ := RoomCollection.CountDocuments(context.TODO(), bson.M{})
	println(studentCount, studentLogCount, visitorCount, visitorLogCount, roomCount)
	stats := models.DashboardStats{
		TotalStrudents: int(studentCount),
		TotalVisitor:   int(visitorCount),
		TotalAdmin:     int(adminCount),
		TotalLogs:      int(studentLogCount + visitorLogCount),
		TotalRooms:     int(roomCount),
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "found..", "data": stats})
}
