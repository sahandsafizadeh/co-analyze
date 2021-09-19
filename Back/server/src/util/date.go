package util

import "time"

func GetTodayDate() time.Time {
	nowTime := time.Now()
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.UTC)
}
