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

var VisitorCollection = config.Visitor

func GetAllVisitors(c *gin.Context) {
	var visitors []models.Visitor

	result, err := VisitorCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	result.Next(context.TODO())
	{
		var singleObject models.Visitor
		err := result.Decode(&singleObject)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "internal server error", "data": err})
			return
		}
		visitors = append(visitors, singleObject)
	}

	if visitors != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "visitors count" + strconv.Itoa(len(visitors)), "data": visitors})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found.."})
}

func RemoveVisitor(c *gin.Context) {
	visitorId := c.Param("visitorId")

	objId, _ := primitive.ObjectIDFromHex(visitorId)

	var user models.Visitor

	deletedUser := VisitorCollection.FindOne(context.TODO(), bson.M{"_id": objId})

	deletedUser.Decode(&user)

	result, err := VisitorCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

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

func SearchVisitor(c *gin.Context) {
	quary := c.Param("quary")

	model := mongo.IndexModel{Keys: bson.D{{Key: "fullname", Value: "text"}, {Key: "city", Value: "text"}}}
	VisitorCollection.Indexes().CreateOne(context.TODO(), model)

	result, err := VisitorCollection.Find(context.TODO(), bson.M{"$text": bson.M{"$search": quary}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}
	var Visitors []models.Visitor

	for result.Next(context.TODO()) {
		var user models.Visitor
		result.Decode(&user)
		Visitors = append(Visitors, user)
	}

	if len(Visitors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "found..", "data": Visitors})
}
