package ptime

import "time"

// StartOfDate ...
func StartOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetUTC())
}
