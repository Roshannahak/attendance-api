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

var AdminCollection = config.GetCollection("admin")

func GetAdmins(c *gin.Context) {
	var admin []models.Admin

	result, err := AdminCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	for result.Next(context.TODO()) {
		var singleAdmin models.Admin

		result.Decode(&singleAdmin)
		admin = append(admin, singleAdmin)
	}

	if admin != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "record count : " + strconv.Itoa(len(admin)), "data": admin})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
}

func AdminRegistration(c *gin.Context) {
	var admin models.Admin

	bindErr := c.BindJSON(&admin)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": bindErr})
		return
	}

	newAdmin := models.Admin{
		Id:         primitive.NewObjectID(),
		EmpId:      admin.EmpId,
		FullName:   admin.FullName,
		Department: admin.Department,
		ContactNo:  admin.ContactNo,
	}

	_, err := AdminCollection.InsertOne(context.TODO(), newAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully registered..", "data": newAdmin})
}
