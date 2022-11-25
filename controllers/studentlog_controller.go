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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var studentLogCollection = config.StudentLog
var studentCheckInCollection = config.StudentCheckIn

func StudentEntry(c *gin.Context) {
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
	studentId, err := primitive.ObjectIDFromHex(entryRequest.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid user id"})
		return
	}
	var user models.Student
	studentCollection.FindOne(context.TODO(), bson.M{"_id": studentId}).Decode(&user)

	if room.Id == roomId {
		//check log available in checked in list
		checkedInUser := studentCheckInCollection.FindOne(context.TODO(), bson.M{"user._id": studentId, "room._id": roomId})
		var entryResponse models.StudentLogs
		checkedInUser.Decode(&entryResponse)
		if entryResponse.Id.IsZero() {

			checkedIn(c, room, user)

		} else {
			//check time out
			if isTimeOut(entryResponse.InTime) {
				studentCheckInCollection.DeleteOne(context.TODO(), bson.M{"_id": entryResponse.Id})
				checkedIn(c, room, user)
			} else {
				//if log available in checked in list
				checkedOut(c, entryResponse.Id)
			}
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid room id"})
		return
	}
}

func checkedIn(c *gin.Context, room models.Room, user models.Student) {
	newstudentlog := models.StudentLogs{
		Id:      primitive.NewObjectID(),
		Student: user,
		Room:    room,
		InTime:  time.Now().Format(time.RFC3339),
	}

	//insert data in logs
	if result, _ := studentLogCollection.InsertOne(context.TODO(), newstudentlog); result.InsertedID != nil {
		//insert data in checkin list
		_, err := studentCheckInCollection.InsertOne(context.TODO(), newstudentlog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "welcome you have successfully checkedIn"})
		return
	}
}

func isTimeOut(inTimeStr string) bool {
	InTime, _ := time.Parse(time.RFC3339, inTimeStr)
	currentTime := time.Now()
	diff := currentTime.Sub(InTime)
	timeDiff := int(diff.Hours())
	if timeDiff > 7 {
		return true
	}
	return false
}

func checkedOut(c *gin.Context, logId primitive.ObjectID) {

	_, errOnDelete := studentCheckInCollection.DeleteOne(context.TODO(), bson.M{"_id": logId})
	if errOnDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": errOnDelete})
		return
	}

	updateLog := bson.M{"outtime": time.Now().Format(time.RFC3339)}

	result, errOnUpdate := studentLogCollection.UpdateOne(context.TODO(), bson.M{"_id": logId}, bson.M{"$set": updateLog})
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

func GetAllStudentLogs(c *gin.Context) {
	var logs []models.StudentLogs

	opts := options.Find().SetProjection(bson.M{"user.contactno": 0, "user.rollno": 0, "room.created": 0})
	result, err := studentLogCollection.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.StudentLogs

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

func GetStudentCheckedInList(c *gin.Context) {
	var logs []models.StudentLogs

	opts := options.Find().SetProjection(bson.M{"user.contactno": 0, "user.rollno": 0, "room.created": 0})
	result, err := studentCheckInCollection.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.StudentLogs

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

func GetLogsByStudentId(c *gin.Context) {
	studentId := c.Param("studentId")

	var logs []models.StudentLogs

	id, _ := primitive.ObjectIDFromHex(studentId)

	opts := options.Find().SetProjection(bson.M{"user.contactno": 0, "user.rollno": 0, "room.created": 0})
	result, err := studentLogCollection.Find(context.TODO(), bson.M{"user._id": id}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.StudentLogs
		result.Decode(&singleLog)
		logs = append(logs, singleLog)
	}

	if logs != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "total count : ", "data": logs})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "not found"})
}
