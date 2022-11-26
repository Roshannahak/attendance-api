package middleware

import "time"

func IsTimeOut(inTimeStr string) bool {
	InTime, _ := time.Parse(time.RFC3339, inTimeStr)
	currentTime := time.Now()
	diff := currentTime.Sub(InTime)
	timeDiff := int(diff.Hours())
	if timeDiff > 7 {
		return true
	}
	return false
}