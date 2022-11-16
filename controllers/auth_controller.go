package controllers

import (
	"attendance_api/models"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserLogin(c *gin.Context) {
	var auth models.Auth

	if err := c.BindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}
	var user models.User
	userCollection.FindOne(context.TODO(), bson.M{"contactno": auth.ContactNo, "rollno": strings.ToLower(auth.RollNo)}).Decode(&user)

	if user.Id.IsZero() {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Invalid cradentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "loggedin successfully..", "data": user})

}

func UserRegistration(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": err})
		return
	}

	rollNo := strings.ToLower(user.RollNo)

	registeredUser, _ := userCollection.CountDocuments(context.TODO(), bson.M{"rollno": rollNo})
	println(registeredUser)
	if registeredUser > 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "msg": "user already registered"})
		return
	}

	newUser := models.User{
		Id:        primitive.NewObjectID(),
		FullName:  user.FullName,
		RollNo:    rollNo,
		Branch:    user.Branch,
		Course:    user.Course,
		Semester:  user.Semester,
		ContactNo: user.ContactNo,
	}

	_, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully registered..", "data": newUser})
}
