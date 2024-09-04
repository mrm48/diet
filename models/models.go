package models

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Dieter struct {

    ID          string `json:"id"`
    Name        string `json:"name"`
    Calories    int `json:"calories"`

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
func GetDieter(context *gin.Context){
    
    var d Dieter

    r := 0

    if err := context.BindJSON(&d); err != nil {
        return
    }

    for _,v := range Dieters {
        if v.Name == d.Name {
            context.IndentedJSON(http.StatusOK, v)
            r = 1
        }
    }

    if r == 0 {
        context.IndentedJSON(http.StatusNotFound, d)
    }

}

func SetDieterCalories(context *gin.Context){

    var d Dieter 

    r := 0

    if err := context.BindJSON(&d); err != nil {
        return
    }

    for k,v := range Dieters {
        if v.Name == d.Name {
            SetCalories(k, d.Calories)
            v.Calories = d.Calories
            r = 1
            context.IndentedJSON(http.StatusOK, v)
        }
    }

    if r == 0 {
        context.IndentedJSON(http.StatusNotFound, d)
    }

}

func SetCalories(k int, c int) {

    Dieters[k].Calories = c

}
