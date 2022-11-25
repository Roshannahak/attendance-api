package controllers

import (
	"attendance_api/config"
	"attendance_api/models"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var studentCollection = config.Student

func StudentLogin(c *gin.Context) {
	var auth models.Auth

	if err := c.BindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}
	var student models.Student
	studentCollection.FindOne(context.TODO(), bson.M{"contactno": auth.ContactNo, "rollno": strings.ToLower(auth.RollNo)}).Decode(&student)

	if student.Id.IsZero() {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Invalid cradentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "loggedin successfully..", "data": student})

}

func StudentRegistration(c *gin.Context) {
	var student models.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": err})
		return
	}

	rollNo := strings.ToLower(student.RollNo)

	registeredStudent, _ := studentCollection.CountDocuments(context.TODO(), bson.M{"rollno": rollNo})
	println(registeredStudent)
	if registeredStudent > 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "msg": "student already registered"})
		return
	}

	newStudent := models.Student{
		Id:        primitive.NewObjectID(),
		FullName:  student.FullName,
		RollNo:    rollNo,
		Branch:    student.Branch,
		Course:    student.Course,
		Semester:  student.Semester,
		ContactNo: student.ContactNo,
		UserType: "STUDENT",
	}

	_, err := studentCollection.InsertOne(context.TODO(), newStudent)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully registered..", "data": newStudent})
}
