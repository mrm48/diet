package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type dieter struct {

    ID string `json:"id"`
    Name string `json:"name"`
    Calories int `json:"calories"`

}

type ndieter struct {

    Name string `json:"name"`

}

type cdieter struct {
    Name string `json:"name"`
    Calories int `json:"calories"`
}

var dieters = []dieter{
    
    {ID: "1", Name: "Matt", Calories: 1600},
    {ID: "2", Name: "Jack", Calories: 1600},

}

func getDieters(context *gin.Context){

    context.IndentedJSON(http.StatusOK, dieters)

}

func addDieter(context *gin.Context){

    var newDieter dieter 

    if err := context.BindJSON(&newDieter); err != nil {
        return
    }

    dieters = append(dieters, newDieter)

    context.IndentedJSON(http.StatusCreated, newDieter)

}

func getDieter(context *gin.Context){
    
    var d ndieter

    r := 0

    if err := context.BindJSON(&d); err != nil {
        return
    }

    for _,v := range dieters {
        if v.Name == d.Name {
            context.IndentedJSON(http.StatusOK, v)
            r = 1
        }
    }

    if r == 0 {
        context.IndentedJSON(http.StatusNotFound, d)
    }

}

func setDieterCalories(context *gin.Context){

    var d cdieter 

    r := 0

    if err := context.BindJSON(&d); err != nil {
        return
    }

    for k,v := range dieters {
        if v.Name == d.Name {
            setCalories(k, d.Calories)
            v.Calories = d.Calories
            r = 1
            context.IndentedJSON(http.StatusOK, v)
        }
    }

    if r == 0 {
        context.IndentedJSON(http.StatusNotFound, d)
    }

}

func setCalories(k int, c int) {

    dieters[k].Calories = c

}

func main() {

    router := gin.Default()

    // all dieters
    router.GET("/dieters", getDieters)
    router.POST("/dieters", addDieter)

    // single dieter
    router.GET("/dieter", getDieter)
    router.POST("/dieter/calories", setDieterCalories)

    // start server
    router.Run("localhost:9090")

}
