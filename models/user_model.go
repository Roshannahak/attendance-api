package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FullName  string             `json:"firstname"`
	RollNo    string             `json:"rollno"`
	Branch    string             `json:"branch"`
	Course    string             `json:"course"`
	Semester  int                `json:"semester"`
	ContactNo string             `json:"contactno"`
}
