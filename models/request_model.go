package models

type EntryLogsRequest struct {
	UserId string `json:"userid"`
	RoomId string `json:"roomid"`
	Reason string `json:"reason"`
}
