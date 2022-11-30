package controllers

import (
	"attendance_api/config"
	"attendance_api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AdminCollection = config.Admin

func AdminRegistration(c *gin.Context) {
	var admin models.Admin

	bindErr := c.BindJSON(&admin)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": bindErr})
		return
	}

	count, _ := AdminCollection.CountDocuments(context.TODO(), bson.M{"empid": admin.EmpId})
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "msg": "user already registered.."})
		return
	}

	newAdmin := models.Admin{
		Id:         primitive.NewObjectID(),
		EmpId:      admin.EmpId,
		FullName:   admin.FullName,
		Department: admin.Department,
		ContactNo:  admin.ContactNo,
		SuperAdmin: admin.SuperAdmin,
	}

	_, err := AdminCollection.InsertOne(context.TODO(), newAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully registered..", "data": newAdmin})
}

func AdminLogin(c *gin.Context) {
	var cradentials models.AdminAuthRequest

	bindErr := c.BindJSON(&cradentials)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "bed request.."})
		return
	}

	var admin models.Admin
	AdminCollection.FindOne(context.TODO(), bson.M{"empid": cradentials.EmpId, "contactno": cradentials.ContactNo}).Decode(&admin)

	if admin.Id.IsZero() {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "invalid cradentials.."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "successfully loggedin..", "data": admin})
}
