package main

import (
    "fmt"
    "diet/models"
    "diet/router"
    "github.com/gin-gonic/gin"
)

var Dieters []models.Dieter

func AddDieters() {

    fmt.Printf("Adding Dieters")
    Dieters = []models.Dieter {
    
        {ID: "1", Name: "Matt", Calories: 1600},
        {ID: "2", Name: "Jack", Calories: 1600},

    }

}

func RunServer() {

    r := gin.Default()

    router.SetRoutes(r)

    // start server
    r.Run("localhost:9099")

}
