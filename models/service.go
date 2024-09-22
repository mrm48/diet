package models

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetDieters(context *gin.Context){

    context.IndentedJSON(http.StatusOK, Dieters)

}

func AddDieter(context *gin.Context){

    var n Dieter 

    if err := context.BindJSON(&n); err != nil {
        return
    }

    Dieters = append(Dieters, n)

    context.IndentedJSON(http.StatusCreated, n)

}

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

func SetCalories(d Dieter, c int) {

    for k, v := range Entries {
        if v.ID == GetID(d) {
            Dieters[k].Calories = c 
        }
    }

}

func GetCalories(d Dieter) int {

    for _, v := range Dieters {
        if v.ID == d.ID {
            return v.Calories
        }
    }

    return 0

}

func GetID(d Dieter) int64 {

    for _, v := range Dieters {
        if v.Name == d.Name {
            return v.ID
        }
    }

    return 0

}

func SetID(d Dieter) {

    for k,v := range Dieters { 
        if v.Name == d.Name {
            Dieters[k].ID = d.ID
        }
    }

}
