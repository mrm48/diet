package models

import (
	"time"
)

type Dieter struct {

    ID          int64       `json:"id"`
    Name        string      `json:"name"`
    Calories    int         `json:"calories"`

}

type Entry struct {

    ID          int64       `json:"id"`
    FoodID      int64       `json:"food"`
    MealID      int64       `json:"meal"`
    Calories    int         `json:"calories"`

}

type Food struct {

    ID          int64       `json:"id"`
    Name        string      `json:"name"`
    Calories    int         `json:"calories"`
    Units       int         `json:"units"`

}

type Meal struct {

    ID          int64       `json:"id"`
    Name        string      `json:"name"`
    Day         time.Time   `json:"day"`
    Calories    int         `json:"calories"`
    Dieter      string      `json:"dieter"`

}

var Dieters []Dieter
var Entries []Entry
var Foods   []Food
var Meals   []Meal

