package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	FullName  string             `json:"fullname,omitempty"`
	RollNo    string             `json:"rollno,omitempty"`
	Branch    string             `json:"branch,omitempty"`
	Course    string             `json:"course,omitempty"`
	Semester  int                `json:"semester,omitempty"`
	ContactNo string             `json:"contactno,omitempty"`
	UserType  string             `json:"usertype,omitempty"`
}
