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
