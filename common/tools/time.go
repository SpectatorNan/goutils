package tools

import "time"

const (
	YYYYMMDDHHmmssNoSplit = "20060102150405"
	YYYYMMDD              = "2006-01-02"
	YYYYMMDDHHmmSS        = "2006-01-02 15:04:05"
	YMDHms                = "2006.01.02 15:04:05"

	YYYYMMDDHHmmSSZone    = "2006-01-02 15:04:04 -0700"
	YYYYMMDDHHmmSSISO8601 = "2006-01-02T15:04:05.000Z"
)

func StringToTime(myTime, format string) (*time.Time, error) {
	parse, err := time.ParseInLocation(myTime, format, time.Local)
	if err != nil {
		return nil, err
	}
	return &parse, nil
}

func TimeToString(myTime time.Time, format string) string {
	return myTime.Format(format)
}

func TimeParseInt64(tstamp int64) *time.Time {
	if tstamp == 0 {
		return nil
	}
	t := time.Unix(tstamp, 0)
	return &t
}
func TimeUnixShowLayoutString(t int64) string {
	return TimeParseInt64(t).Format(YYYYMMDDHHmmSS)
}
func TimeShowLayoutString(t time.Time) string {
	return t.Format(YYYYMMDDHHmmSS)
}

func TimePushDayToString(t time.Time, day int) string {
	return t.AddDate(0, 0, day).Format(YYYYMMDDHHmmSS)
}

func DateTimeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
