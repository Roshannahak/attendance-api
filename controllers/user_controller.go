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
)

var userCollection = config.GetCollection("users")

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": err})
		return
	}

	newUser := models.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		College:   user.College,
		Branch:    user.Branch,
		Course:    user.Course,
		Semester:  user.Semester,
		ContactNo: user.ContactNo,
	}

	result, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
		return
	}

	if result.InsertedID != nil {
		c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully inserted..", "data": newUser})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found.."})
}

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

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "users count" + strconv.Itoa(len(users)), "data": users})
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
