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

var studentCollection = config.Student

func GetAllStudents(c *gin.Context) {
	var students []models.Student

	result, err := studentCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleUser models.Student
		err := result.Decode(&singleUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
			return
		}

		students = append(students, singleUser)
	}
	if students != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "students count" + strconv.Itoa(len(students)), "data": students})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found.."})
}

func RemoveStudent(c *gin.Context) {
	studentId := c.Param("studentId")

	objId, _ := primitive.ObjectIDFromHex(studentId)

	var user models.Student

	deletedUser := studentCollection.FindOne(context.TODO(), bson.M{"_id": objId})

	deletedUser.Decode(&user)

	result, err := studentCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

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

func SearchStudent(c *gin.Context) {
	quary := c.Param("quary")

	model := mongo.IndexModel{Keys: bson.D{{Key: "branch", Value: "text"}, {Key: "fullname", Value: "text"}, {Key: "semester", Value: "text"}, {Key: "course", Value: "text"}}}
	studentCollection.Indexes().CreateOne(context.TODO(), model)

	result, err := studentCollection.Find(context.TODO(), bson.M{"$text": bson.M{"$search": quary}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}
	var students []models.Student

	for result.Next(context.TODO()) {
		var user models.Student
		result.Decode(&user)
		students = append(students, user)
	}

	if len(students) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "found..", "data": students})
}
