package mutils

import (
	"time"
)

// GetCurrentDate Returns the current date as a string for storing meals/entries by day
func GetCurrentDate() string {
	day := time.Now().Format("2006-01-02T15:04:05 -070000")

	day = day[:10]

	return day
}

// GetDateFromString will parse a time.Time object from a string that can be used to query the database for a specific day's meals/entries
func GetDateFromString(date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02T15:04:05 -070000", date)
	return parsedDate, err
}
