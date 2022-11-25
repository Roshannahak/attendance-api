package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentLogs struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Student Student            `json:"student"`
	Room    Room               `json:"room"`
	InTime  string             `json:"intime"`
	OutTime string             `json:"outtime"`
}
