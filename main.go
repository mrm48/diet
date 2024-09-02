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

func main() {
    router := gin.Default()
    router.GET("/dieters", getDieters)
    router.POST("/dieters", addDieter)
    router.Run("localhost:9090")
}
