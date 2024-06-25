package ptime

import "time"

// Now ...
func Now() time.Time {
	return time.Now().In(GetUTC())
}
