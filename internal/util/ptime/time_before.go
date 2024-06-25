package ptime

import "time"

// TimeBeforeNowInMin ...
func TimeBeforeNowInMin(min int) time.Time {
	return time.Now().Add(time.Minute * time.Duration(min) * -1).In(GetUTC())
}
