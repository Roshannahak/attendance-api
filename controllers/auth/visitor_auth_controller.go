package controllers

import (
	"attendance_api/config"
	"attendance_api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var visitorCollaction = config.Visitor

func VisitorLogin(c *gin.Context) {
	var visitor models.Visitor

	if err := c.BindJSON(&visitor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	newVisitor := models.Visitor{
		FullName:  visitor.FullName,
		City:      visitor.City,
		ContactNo: visitor.ContactNo,
		UserType: "VISITOR",
	}

	_, err := visitorCollaction.InsertOne(context.TODO(), newVisitor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "msg": "successfully logedin..", "data": newVisitor})
}
