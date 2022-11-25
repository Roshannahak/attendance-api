package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	EmpId      string             `json:"empid,omitempty"`
	FullName   string             `json:"fullname,omitempty"`
	Department string             `json:"department,omitempty"`
	ContactNo  string             `json:"contactno,omitempty"`
	SuperAdmin bool               `json:"superadmin,omitempty"`
}
