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
        context.IndentedJSON(http.StatusNotFound, nil)
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
        context.IndentedJSON(http.StatusNotFound, nil)
    }

}

func SetCalories(k int, c int) {

    Dieters[k].Calories = c

}
