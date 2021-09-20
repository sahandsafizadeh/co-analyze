package util

import "time"

func GetTodayDate() time.Time {
	nowTime := time.Now()
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.UTC)
}

func GetYesterdayDate() time.Time {
	today := GetTodayDate()
	return today.Add(-24 * time.Hour)
}
