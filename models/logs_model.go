package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EntryLogsResponse struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	User    User               `json:"user"`
	Room    Room               `json:"room"`
	InTime  string             `json:"intime"`
	OutTime string             `json:"outtime"`
}

type EntryLogsRequest struct {
	UserId string `json:"userid"`
	RoomId string `json:"roomid"`
}
