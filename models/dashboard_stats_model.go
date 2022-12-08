package models

type DashboardStats struct {
	TotalStrudents int `json:"total_student"`
	TotalVisitor   int `json:"total_visitor"`
	TotalAdmin     int `json:"total_admin"`
	TotalLogs      int `json:"total_logs"`
	TotalRooms     int `json:"total_rooms"`
}
