package models

type AdminAuthRequest struct {
	EmpId     string `json:"empid,omitempty"`
	ContactNo string `json:"contactno,omitempty"`
}
