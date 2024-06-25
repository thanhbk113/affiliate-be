package ptime

import "time"

// FromUnix ...
func FromUnix(s int64) time.Time {
	return time.Unix(s, 0)
}

// TimeParseISODate ...
func TimeParseISODate(value string) time.Time {
	t, _ := time.Parse(DateLayoutFull, value)
	return t
}

// TimeOfDayInHCM ...
func TimeOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(timezoneHCM)
	return t.In(l)
}

// TimeStartOfDayInHCM ...
func TimeStartOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(timezoneHCM)
	y, m, d := t.In(l).Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, l).UTC()
	return date
}

// ParseWithUTC ...
func ParseWithUTC(s, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}
