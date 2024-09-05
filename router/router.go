package router

import (
	"mauit/models"

	"github.com/gin-gonic/gin"
)

func SetRoutes(Router *gin.Engine) {

    // all dieters
    Router.GET("/dieters", models.GetDieters)
    Router.POST("/dieters", models.AddDieter)

    // single dieter
    Router.GET("/dieter", models.GetDieter)
    Router.POST("/dieter/calories", models.SetDieterCalories)

}
