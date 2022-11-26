package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Visitor struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	FullName  string             `json:"fullname,omitempty"`
	City      string             `json:"city,omitempty"`
	ContactNo string             `json:"contactno,omitempty"`
	UserType  string             `json:"usertype,omitempty"`
}
