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

var logCollection = config.GetCollection("logs")
var checkInCollection = config.GetCollection("checkin")

func CheckedIn(c *gin.Context) {
	var userLog models.Logs

	if err := c.BindJSON(&userLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}

	roomId, err := primitive.ObjectIDFromHex(userLog.RoomId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid room id"})
		return
	}

	//validate room id
	var room models.Room
	RoomCollection.FindOne(context.TODO(), bson.M{"_id": roomId}).Decode(&room)
	if room.Id == roomId {
		//check log available in checked in list
		checkedInUser := checkInCollection.FindOne(context.TODO(), bson.M{"userid": userLog.UserId, "roomid": userLog.RoomId})
		var getLog models.Logs
		checkedInUser.Decode(&getLog)
		if getLog.Id.IsZero() {

			// println(getLog.Id.String())

			// //check time out
			// if isTimeOut(getLog.InTime){
			// 	checkInCollection.DeleteOne(context.TODO(), bson.M{"_id": getLog.Id})
			// }

			newlog := models.Logs{
				Id:     primitive.NewObjectID(),
				UserId: userLog.UserId,
				RoomId: userLog.RoomId,
				InTime: time.Now().Format(time.RFC3339),
			}

			//insert data in logs
			if result, _ := logCollection.InsertOne(context.TODO(), newlog); result.InsertedID != nil {
				//insert data in checkin list
				_, err := checkInCollection.InsertOne(context.TODO(), newlog)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
					return
				}
				c.JSON(http.StatusOK, gin.H{"success": true, "msg": "welcome you have successfully checkedIn"})
				return
			}

		} else {
			//if log available in checked in list
			checkedOut(c, getLog.Id)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid room id"})
		return
	}
}

func isTimeOut(inTimeStr string) bool {
	InTime, _ := time.Parse(time.RFC3339, inTimeStr)
	currentTime := time.Now()
	diff := currentTime.Sub(InTime)
	timeDiff := int(diff.Hours())
	if timeDiff > 12 {
		println(true)
		return true
	}
	println(false)
	return false
}

func checkedOut(c *gin.Context, logId primitive.ObjectID) {

	_, errOnDelete := checkInCollection.DeleteOne(context.TODO(), bson.M{"_id": logId})
	if errOnDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": errOnDelete})
		return
	}

	updateLog := bson.M{"outtime": time.Now().Format(time.RFC3339)}

	result, errOnUpdate := logCollection.UpdateOne(context.TODO(), bson.M{"_id": logId}, bson.M{"$set": updateLog})
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

func GetAllLogs(c *gin.Context) {
	var logs []models.Logs

	result, err := logCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.Logs

		result.Decode(&singleLog)
		logs = append(logs, singleLog)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "userLog count : " + strconv.Itoa(len(logs)), "data": logs})
}

func GetCheckedInList(c *gin.Context) {
	var logs []models.Logs

	result, err := checkInCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleLog models.Logs

		result.Decode(&singleLog)
		logs = append(logs, singleLog)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "checked in count : " + strconv.Itoa(len(logs)), "data": logs})
}
