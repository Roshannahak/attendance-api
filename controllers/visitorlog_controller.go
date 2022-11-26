package controllers

import (
	"attendance_api/config"
	"attendance_api/middleware"
	"attendance_api/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var visitorLogCollection = config.VisitorLog
var visitorCheckInCollection = config.VisitorCheckIn
var visitorCollection = config.Visitor

func VisitorEntry(c *gin.Context) {
	var entryRequest models.EntryLogsRequest

	if err := c.BindJSON(&entryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}

	//find room object
	roomId, err := primitive.ObjectIDFromHex(entryRequest.RoomId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid room id"})
		return
	}
	var room models.Room
	RoomCollection.FindOne(context.TODO(), bson.M{"_id": roomId}).Decode(&room)

	//find user object
	visitorId, err := primitive.ObjectIDFromHex(entryRequest.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid user id"})
		return
	}
	var user models.Visitor
	visitorCheckInCollection.FindOne(context.TODO(), bson.M{"_id": visitorId}).Decode(&user)

	if room.Id == roomId {
		//check log available in checked in list
		checkedInUser := visitorCheckInCollection.FindOne(context.TODO(), bson.M{"visitor._id": visitorId, "room._id": roomId})
		var entryResponse models.VisitorLogs
		checkedInUser.Decode(&entryResponse)
		if entryResponse.Id.IsZero() {

			visitorCheckedIn(c, room, user)

		} else {
			//check time out
			if middleware.IsTimeOut(entryResponse.InTime) {
				visitorCheckInCollection.DeleteOne(context.TODO(), bson.M{"_id": entryResponse.Id})
				visitorCheckedIn(c, room, user)
			} else {
				//if log available in checked in list
				visitorCheckedOut(c, entryResponse.Id)
			}
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid room id"})
		return
	}
}

func visitorCheckedIn(c *gin.Context, room models.Room, user models.Visitor) {
	newVisitorlog := models.VisitorLogs{
		Id:      primitive.NewObjectID(),
		Visitor: user,
		Room:    room,
		InTime:  time.Now().Format(time.RFC3339),
	}

	//insert data in logs
	if result, _ := visitorLogCollection.InsertOne(context.TODO(), newVisitorlog); result.InsertedID != nil {
		//insert data in checkin list
		_, err := visitorCheckInCollection.InsertOne(context.TODO(), newVisitorlog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "welcome you have successfully visitorCheckedIn"})
		return
	}
}

func visitorCheckedOut(c *gin.Context, logId primitive.ObjectID) {

	_, errOnDelete := visitorCheckInCollection.DeleteOne(context.TODO(), bson.M{"_id": logId})
	if errOnDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": errOnDelete})
		return
	}

	updateLog := bson.M{"outtime": time.Now().Format(time.RFC3339)}

	result, errOnUpdate := visitorLogCollection.UpdateOne(context.TODO(), bson.M{"_id": logId}, bson.M{"$set": updateLog})
	if errOnUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": errOnUpdate})
		return
	}
	if result.MatchedCount == 1 {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "checked out successfully.."})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found.."})
}

func GetAllVisitorLogs(c *gin.Context) {
	var logs []models.VisitorLogs

	opts := options.Find().SetProjection(bson.M{"user.contactno": 0, "user.rollno": 0, "room.created": 0})
	result, err := visitorLogCollection.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.VisitorLogs

		result.Decode(&singleLog)
		logs = append(logs, singleLog)
	}
	if logs != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "userLog count : " + strconv.Itoa(len(logs)), "data": logs})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "not found"})
		return
	}
}

func GetVisitorCheckedInList(c *gin.Context) {
	var logs []models.VisitorLogs

	opts := options.Find().SetProjection(bson.M{"visitor.contactno": 0, "visitor.rollno": 0, "room.created": 0})
	result, err := visitorCheckInCollection.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.VisitorLogs

		result.Decode(&singleLog)
		logs = append(logs, singleLog)
	}
	if logs != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "checked in count : " + strconv.Itoa(len(logs)), "data": logs})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "not found"})
		return
	}
}

func GetLogsByVisitorId(c *gin.Context) {
	visitorId := c.Param("visitorId")

	var logs []models.VisitorLogs

	id, _ := primitive.ObjectIDFromHex(visitorId)

	opts := options.Find().SetProjection(bson.M{"visitor.contactno": 0, "visitor.rollno": 0, "room.created": 0})
	result, err := visitorLogCollection.Find(context.TODO(), bson.M{"visitor._id": id}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.VisitorLogs
		result.Decode(&singleLog)
		logs = append(logs, singleLog)
	}

	if logs != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "total count : ", "data": logs})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "not found"})
}
