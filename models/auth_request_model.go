package models

type StudentAuthRequest struct{
	ContactNo string             `json:"contactno"`
	RollNo    string             `json:"password"`
}

type AdminAuthRequest struct {
	EmpId     string `json:"empid,omitempty"`
	ContactNo string `json:"contactno,omitempty"`
}