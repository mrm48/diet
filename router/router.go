package router

import (
	"mauit/models"

	"github.com/gin-gonic/gin"
)

func SetRoutes(Router *gin.Engine) {

    // all dieters
    Router.GET("/dieters/all", models.GetDieters)
    Router.POST("/dieters/add", models.AddDieter)

    // single dieter
    Router.POST("/dieter/search/name", models.GetDieter)
    Router.POST("/dieter/set/calories", models.SetDieterCalories)

}
