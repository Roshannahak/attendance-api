package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Logs struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	UserId  string             `json:"userid"`
	RoomId  string             `json:"roomid"`
	InTime  string             `json:"intime"`
	OutTime string             `json:"outtime"`
}
