package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	College   string             `json:"college"`
	Branch    string             `json:"branch"`
	Course    string             `json:"course"`
	Semester  int                `json:"semester"`
	ContactNo string             `json:"contactno"`
}
