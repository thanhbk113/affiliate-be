package ptime

import (
	"encoding/json"
	"time"
)

// TimeResponse ...
type TimeResponse struct {
	Time time.Time
}

// UnmarshalJSON ...
func (t *TimeResponse) UnmarshalJSON(b []byte) error {
	if string(b) == "" || string(b) == "\"\"" {
		return nil
	}
	return json.Unmarshal(b, &t.Time)
}

// MarshalJSON ...
func (t *TimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.FormatISODate())
}

// FormatISODate ...
func (t *TimeResponse) FormatISODate() string {
	if t == nil || t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(DateLayoutFull)
}

// TimeResponseInit ...
func TimeResponseInit(t time.Time) *TimeResponse {
	return &TimeResponse{Time: t}
}
