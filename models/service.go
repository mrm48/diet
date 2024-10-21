package models

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetDieters(context *gin.Context){

    context.IndentedJSON(http.StatusOK, Dieters)

}

// Add specifically a dieter
func AddDieter(context *gin.Context){

    var n Dieter 

    if err := context.BindJSON(&n); err != nil {
        return
    }

    Dieters = append(Dieters, n)

    context.IndentedJSON(http.StatusCreated, n)

}

// Get dieter by name
func GetDieter(context *gin.Context){
    
    var d Dieter

    if err := context.BindJSON(&d); err != nil {
        return
    }

    for _,v := range Dieters {
        if v.Name == d.Name {
            context.IndentedJSON(http.StatusOK, v)
            return
        }
    }

    context.IndentedJSON(http.StatusNotFound, nil)

}

// Set the calories available for a dieter
func SetDieterCalories(context *gin.Context){

    var d Dieter 

    if err := context.BindJSON(&d); err != nil {
        return
    }

    for _,v := range Dieters {
        if v.Name == d.Name {
            SetCalories(v, d.Calories)
            v.Calories = d.Calories
            context.IndentedJSON(http.StatusOK, v)
            return
        }
    }

    context.IndentedJSON(http.StatusNotFound, nil)

}

// Set the saved dieter's number of maximum calories
func SetCalories(d Dieter, c int) {

    for k, v := range Entries {
        if v.ID == GetID(d) {
            Dieters[k].Calories = c 
        }
    }

}

// Get maximum number of calories for a dieter
func GetCalories(d Dieter) int {

    for _, v := range Dieters {
        if v.ID == d.ID {
            return v.Calories
        }
    }

    return 0

}

// Get the unique ID for the dieter by name
func GetID(d Dieter) int64 {

    for _, v := range Dieters {
        if v.Name == d.Name {
            return v.ID
        }
    }

    return 0

}

// Set the unique ID for a dieter by name
func SetID(d Dieter) {

    for k,v := range Dieters { 
        if v.Name == d.Name {
            Dieters[k].ID = d.ID
        }
    }

}
