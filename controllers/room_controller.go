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

var roomCollection = config.GetCollection("Rooms")

func CreateRoom(c *gin.Context) {
	var room models.Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}

	newRoom := models.Room{
		Id:             primitive.NewObjectID(),
		RoomNo:         room.RoomNo,
		RoomId:         room.RoomId,
		DepartmentName: room.DepartmentName,
		CreatedAt:      time.Now().Local(),
	}

	_, err := roomCollection.InsertOne(context.TODO(), newRoom)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully created..", "data": newRoom})
}

func GetAllRooms(c *gin.Context) {
	var rooms []models.Room

	result, err := roomCollection.Find(context.TODO(), bson.M{})

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
