package models

import (
	"time"
)

// Describe a record in the database
type Record interface {
	GetID() int64
	SetID(int64)
}

// Create an interface for items that have a calories property
type Named interface {
	GetName() string
	SetName(string)
	GetCalories() int
	SetCalories(int)
}

// User struct, maximum daily calories and name
type Dieter struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Calories int    `json:"calories"`
}

// Describe an item in the database for multiple food items associated to one meal
type Entry struct {
	ID       int64 `json:"id"`
	FoodID   int64 `json:"food"`
	MealID   int64 `json:"meal"`
	Calories int   `json:"calories"`
}

// Describe a food that can be added to a meal by a user
type Food struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Calories int    `json:"calories"`
	Units    int    `json:"units"`
}

// Describe a meal that can be consumed by a user
type Meal struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Day      string `json:"day"`
	Calories int    `json:"calories"`
	Dieterid int64  `json:"dieterid"`
	Dieter   string `json:"dieter"`
}

// GetCurrentDate Returns the current date as a string for storing meals/entries by day
func GetCurrentDate() string {
	day := time.Now().Format("2006-01-02T15:04:05 -070000")

	day = day[:10]

	return day
}
