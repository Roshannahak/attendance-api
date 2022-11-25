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

var RoomCollection = config.Room

func CreateRoom(c *gin.Context) {
	var room models.Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}
	newRoom := models.Room{
		Id:             primitive.NewObjectID(),
		RoomNo:         room.RoomNo,
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

func DeleteRoom(c *gin.Context) {
	roomId := c.Param("roomId")

	objId, _ := primitive.ObjectIDFromHex(roomId)

	result, err := RoomCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	if result.DeletedCount == 1 {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "room successfully deleted.."})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "room not found.."})
}

func UpdateRoom(c *gin.Context) {
	roomId := c.Param("roomId")

	objId, _ := primitive.ObjectIDFromHex(roomId)

	var room models.Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	updated := bson.M{"roomno": room.RoomNo, "departmentname": room.DepartmentName}

	updateResult, err := RoomCollection.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": updated})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	if updateResult.ModifiedCount == 1{
		RoomCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&room)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "successfully updated..", "data": room})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found.."})
}
