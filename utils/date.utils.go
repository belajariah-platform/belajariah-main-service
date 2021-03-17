package utils

import (
	"fmt"
	"time"
)

func CurrentDateString(date time.Time) string {
	dateString := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.000Z",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())
	return dateString
}

func CurrentDateStringCustom(date time.Time) string {
	dateString := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())
	return dateString
}

func CurrentDate(date time.Time) string {
	dateString := fmt.Sprintf("%d-%02d-%02d",
		date.Year(), date.Month(), date.Day())
	return dateString
}

func CurrentTime(date time.Time) string {
	dateString := fmt.Sprintf("%02d:%02d:%02d",
		date.Hour(), date.Minute(), date.Second())
	return dateString
}

func HandleNullableStringDate(date string) string {
	if len(date) > 0 {
		stringDate := "'" + date + "'"
		return stringDate
	} else {
		return "null"
	}
}

func HandleAddDate(date time.Time, status string) string {
	var dates string
	if status == "Waiting for Payment" {
		return HandleNullableDate(date.Add(24 * time.Hour))
	}
	return dates
}

func HandleNullableDate(date time.Time) string {
	if date.IsZero() {
		return ""
	} else {
		stringDate := CurrentDateString(date)
		return stringDate
	}
}

func GetDuration(date1, date2 time.Time) float64 {
	duration := date2.Sub(date1).Minutes()
	return duration
}

func TimeAdd(date1 time.Time, duration int) time.Time {
	date2 := date1.AddDate(0, 0, duration*30)
	return date2
}
