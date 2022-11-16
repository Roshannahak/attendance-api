package controllers

import (
	"attendance_api/config"
	"attendance_api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = config.GetCollection("users")

func GetAllUser(c *gin.Context) {
	var users []models.User

	result, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleUser models.User
		err := result.Decode(&singleUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
			return
		}

		users = append(users, singleUser)
	}
	if users != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "users count" + strconv.Itoa(len(users)), "data": users})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "users not found.."})
}

func RemoveUser(c *gin.Context) {
	userId := c.Param("userId")

	objId, _ := primitive.ObjectIDFromHex(userId)

	var user models.User

	deletedUser := userCollection.FindOne(context.TODO(), bson.M{"_id": objId})

	deletedUser.Decode(&user)

	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
		return
	}

	if result.DeletedCount == 1 {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "successfully deleted..", "data": user})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found.."})
}

func SearchUser(c *gin.Context) {
	quary := c.Param("quary")

	model := mongo.IndexModel{Keys: bson.D{{Key: "branch", Value: "text"}, {Key: "fullname", Value: "text"}, {Key: "semester", Value: "text"}, {Key: "course", Value: "text"}}}
	userCollection.Indexes().CreateOne(context.TODO(), model)

	result, err := userCollection.Find(context.TODO(), bson.M{"$text": bson.M{"$search": quary}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}
	var users []models.User

	for result.Next(context.TODO()) {
		var user models.User
		result.Decode(&user)
		users = append(users, user)
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "found..", "data": users})
}
