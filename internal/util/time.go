package util

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const DateISOFormat = "2006-01-02T15:04:05.000Z"
const TimezoneHCM = "Asia/Ho_Chi_Minh"

// FormatTime ...
const (
	FormatTimeExcel         = "2006-01-02 15:04:05"
	FormatTimeDob           = "02/01/2006"
	FormatTimeSendMailBuyer = "15:04 02-01-2006"
	TimeLayoutNoHour        = "02/01/2006"
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

// Format tine ISO
func TimeISO(t time.Time) string {
	return t.Format(DateISOFormat)
}

// MarshalJSON ...
func (t TimeResponse) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(t.Time.Format(DateISOFormat))
}

// TimeResponseInit ...
func TimeResponseInit(t time.Time) TimeResponse {
	return TimeResponse{Time: t}
}

// TimeResponsePointInit ...
func TimeResponsePointInit(t time.Time) *TimeResponse {
	return &TimeResponse{Time: t}
}

var TimeLocationHCM, _ = time.LoadLocation(TimezoneHCM)

// TimeParseWithHCMLocation ...
func TimeParseWithHCMLocation(s, layout string) time.Time {
	t, err := time.ParseInLocation(layout, s, TimeLocationHCM)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t
}

// TimeParseISODate ...
func TimeParseISODate(value string) time.Time {
	t, _ := time.Parse(DateISOFormat, value)
	return t
}

// GetCustomTimeHCMString ...
func GetCustomTimeHCMString(t time.Time, format string) string {
	l, _ := time.LoadLocation(TimezoneHCM)
	y, m, d := t.In(l).Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, l)
	return date.Format(format)
}

// GetDayInWeek
func GetDayInWeek(w time.Weekday) int {
	switch w {
	case time.Monday:
		return 0
	case time.Tuesday:
		return 1
	case time.Wednesday:
		return 2
	case time.Thursday:
		return 3
	case time.Friday:
		return 4
	case time.Saturday:
		return 5
	case time.Sunday:
		return 6
	}
	return 0
}

// TimeOfDayInHCM ...
func TimeOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	return t.In(l)
}

// GetDateInHCM ...
func GetDateInHCM(t time.Time) (year int, month int, day int) {
	y, m, d := TimeOfDayInHCM(t).Date()
	return y, int(m), d
}

// GetStartEndDayOfMonthByTime ...
func GetStartEndDayOfMonthByTime(deliveredAt time.Time) (time.Time, time.Time) {
	t := TimeOfDayInHCM(deliveredAt)
	l, _ := time.LoadLocation(TimezoneHCM)
	y, m, _ := t.Date()
	from := time.Date(y, m, 1, 0, 0, 0, 0, l)
	to := from.AddDate(0, 1, 0)
	return from, to
}

// TimeEndDayNextMonthInHCM ...
func TimeEndDayNextMonthInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	timeHCM := t.In(l)
	y, m, _ := timeHCM.Date()
	return time.Date(y, m, 1, 23, 59, 59, 599, l).AddDate(0, 2, -1)
}

// TimeStartDayMonthInHCM ...
func TimeStartDayMonthInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	timeHCM := t.In(l)
	y, m, _ := timeHCM.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, l)
}

// TimeStartDayMonth ...
func TimeStartDayMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// TimeEndDayMonthInHCM ...
func TimeEndDayMonthInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	timeHCM := t.In(l)
	y, m, _ := timeHCM.Date()
	return time.Date(y, m, 1, 23, 59, 59, 599, l).AddDate(0, 1, -1)
}

// TimeStartOfDayInHCM ...
func TimeStartOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	y, m, d := t.In(l).Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, l).UTC()
	return date
}

// TimeStartOfDayInHCMByDate ...
func TimeStartOfDayInHCMByDate(year int, month time.Month, day int) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	date := time.Date(year, month, day, 0, 0, 0, 0, l).UTC()
	return date
}

// TimeEndOfDayInHCM ...
func TimeEndOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	y, m, d := t.In(l).Date()
	date := time.Date(y, m, d, 23, 59, 59, 599, l).UTC()
	return date
}

// TimeStartOfDay ...
func TimeStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GenerateTime ...
func GenerateTime(format string) string {
	return time.Now().Format(format)
}

func timeInRangeCondition() bson.M {
	now := time.Now()

	return bson.M{
		"$and": []bson.M{
			{
				"$or": []bson.M{
					{"startAt": bson.M{"$lt": now}},
					{"startAt": bson.M{"$exists": false}},
				},
			},
			{
				"$or": []bson.M{
					{"endAt": bson.M{"$gt": now}},
					{"endAt": bson.M{"$exists": false}},
				},
			},
		},
	}
}

// TimeAssignRangeCondition ...
func TimeAssignRangeCondition(cond bson.M) bson.M {
	rangeCond := timeInRangeCondition()

	for k, v := range rangeCond {
		cond[k] = v
	}
	return cond
}

// GetFirstAndLastOfMonth ...
func GetFirstAndLastOfMonth(t time.Time) (firstOfMonth time.Time, lastOfMonth time.Time) {
	currentYear, currentMonth, _ := t.Date()
	l, _ := time.LoadLocation(TimezoneHCM)
	firstOfMonth = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, l)
	y, m, d := firstOfMonth.AddDate(0, 1, -1).In(l).Date()
	lastOfMonth = time.Date(y, m, d, 23, 59, 59, 0, l)
	return
}

// GetLastOfNextMonth ...
func GetLastOfNextMonth() (result time.Time) {
	_, lastMonth := GetFirstAndLastOfMonth(time.Now())
	_, result = GetFirstAndLastOfMonth(lastMonth.AddDate(0, 0, 2))
	return
}

// TimeFormatDob ...
func TimeFormatDob(date time.Time) string {
	l, _ := time.LoadLocation(TimezoneHCM)
	return date.In(l).Format(FormatTimeDob)
}

// SubtractTime ...
func SubtractTime(endAt time.Time) (timeText string) {
	now := time.Now()
	if endAt.Before(now) {
		return
	}

	subHour := endAt.Sub(now).Hours()
	if subHour > 24 {
		mod := int(subHour) % 24

		if mod >= 12 {
			timeText = fmt.Sprintf("Còn %d ngày", int(subHour/24+1))
		} else {
			timeText = fmt.Sprintf("Còn %d ngày", int(subHour/24))
		}
		return
	}

	if subHour >= 1 {
		timeText = fmt.Sprintf("Còn %d giờ", int(subHour))
		return
	}

	minutes := endAt.Sub(now).Minutes()
	if minutes >= 1 && minutes <= 60 {
		timeText = fmt.Sprintf("Còn %d phút", int(minutes))
	} else {
		seconds := endAt.Sub(now).Seconds()
		timeText = fmt.Sprintf("Còn %d giây", int(seconds))
	}

	return
}

// TimeFromUnix ...
func TimeFromUnix(u int64) time.Time {
	t := time.Unix(u, 0)
	return t
}

// TimeFormatInHCMLocation ...
func TimeFormatInHCMLocation(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	}
	return t.In(TimeLocationHCM).Format(layout)
}
