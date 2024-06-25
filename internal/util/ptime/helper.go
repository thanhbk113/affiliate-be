package ptime

import (
	"time"
)

//
// NOTE: due to unique timezone in server's code, all using time will be converted to UTC timezone (UTC +0)
// All functions generate time, must be call util functions here
// WARNING: don't accept call time.Now() directly
//

// GetUTC ...
func GetUTC() *time.Location {
	l, _ := time.LoadLocation("UTC")
	return l
}
