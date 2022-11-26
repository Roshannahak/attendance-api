package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VisitorLogs struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Visitor Visitor            `json:"visitor"`
	Room    Room               `json:"room"`
	InTime  string             `json:"intime"`
	OutTime string             `json:"outtime"`
	Reason  string             `json:"reason"`
}
