package timex

import "time"

const FORMAT = "2006-01-02 15:04:05"

func DateToString(date time.Time) string {
	return date.Format(FORMAT)
}

func StrToDate(date string) time.Time {
	res, _ := time.Parse(FORMAT, date)
	return res
}
