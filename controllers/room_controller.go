package controllers

import (
	"attendance_api/config"
	"attendance_api/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var RoomCollection = config.GetCollection("Rooms")

func CreateRoom(c *gin.Context) {
	var room models.Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}
	roomid := "RM" + room.RoomNo
	newRoom := models.Room{
		Id:             primitive.NewObjectID(),
		RoomNo:         room.RoomNo,
		RoomId:         roomid,
		DepartmentName: room.DepartmentName,
		Created:        time.Now().Format("2006-01-02"),
	}

	_, err := RoomCollection.InsertOne(context.TODO(), newRoom)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully created..", "data": newRoom})
}

func GetAllRooms(c *gin.Context) {
	var rooms []models.Room

	result, err := RoomCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleRoom models.Room

		err := result.Decode(&singleRoom)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
			return
		}

		rooms = append(rooms, singleRoom)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "rooms count : " + strconv.Itoa(len(rooms)), "data": rooms})
}
