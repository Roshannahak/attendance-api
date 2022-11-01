package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	RoomId         string             `json:"roomId"`
	RoomNo         string             `json:"roomNo"`
	DepartmentName string             `json:"departmentName"`
	CreatedAt      time.Time          `json:"cteatedAt"`
}
