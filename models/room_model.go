package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	RoomId         string             `json:"roomid"`
	RoomNo         string             `json:"roomno"`
	DepartmentName string             `json:"departmentname"`
	Created        string             `json:"created"`
}
