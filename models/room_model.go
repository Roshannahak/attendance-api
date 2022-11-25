package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	RoomNo         string             `json:"roomno,omitempty"`
	RoomName       string             `json:"roomname,omitempty"`
	DepartmentName string             `json:"departmentname,omitempty"`
	Created        string             `json:"created,omitempty"`
}
