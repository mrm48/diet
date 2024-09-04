package models

import (
	"time"
)

type Dieter struct {

    ID          string      `json:"id"`
    Name        string      `json:"name"`
    Calories    int         `json:"calories"`

}

type Entry struct {

    ID          string      `json:"id"`
    FoodID      string      `json:"food"`
    MealID      string      `json:"meal"`
    Calories    int         `json:"calories"`

}

type Food struct {

    ID          int         `json:"id"`
    Name        string      `json:"name"`
    Calories    int         `json:"calories"`
    Units       int         `json:"units"`

}

type Meal struct {

    ID          int         `json:"id"`
    Name        string      `json:"name"`
    Day         time.Time   `json:"day"`
    Calories    int         `json:"calories"`
    Dieter      string      `json:"dieter"`

}

var Dieters []Dieter
var Entries []Entry
var Foods   []Food
var Meals   []Meal

