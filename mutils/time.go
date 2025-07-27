package mutils

import (
	"time"
)

var dateTimeStamp string = "2006-01-02T15:04:05 -070000"
var dateStamp string = "2006-01-02"

// GetCurrentDate Returns the current date as a string for storing meals/entries by day
func GetCurrentDate() string {
	day := time.Now().Format(dateTimeStamp)

	day = day[:10]

	return day
}

// GetDateFromString will parse a time.Time object from a string and format it correctly so that it can be used to query the database for a specific day's meals/entries
func GetDateFromString(date string) (string, error) {
	if date != "" {
		parsedDate, err := time.Parse(dateTimeStamp, date)
		if err != nil {
			return "", err
		}
		stringDate := parsedDate.Format(dateStamp)
		return stringDate, err
	} else {
		return time.Now().Format(dateStamp), nil
	}
}
