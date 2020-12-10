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

func HandleNullableStringDate(date string) string {
	if len(date) > 0 {
		stringDate := "'" + date + "'"
		return stringDate
	} else {
		return "null"
	}
}

func HandleNullableDate(date time.Time) string {
	if date.IsZero() {
		return ""
	} else {
		stringDate := CurrentDateString(date)
		return stringDate
	}
}
