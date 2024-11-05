package router

import (
	"mauit/models"

	"github.com/gin-gonic/gin"
)

// Add endpoints to the API, noun verb object pathing
func SetRoutes(Router *gin.Engine) {

	// all dieters
	Router.GET("/dieters/all", models.GetDieters)
	Router.POST("/dieters/add", models.AddDieter)

	// single dieter
	Router.GET("/dieter/search/name", models.GetDieter)
	Router.POST("/dieter/set/calories", models.SetDieterCalories)
	Router.GET("/dieter/search/calories", models.GetDieterCalories)
	Router.GET("/dieter/check/calories", models.GetRemainingDieterCalories)

	// meal
	Router.GET("/meal/search", models.GetMeal)
	Router.POST("/meal/entry", models.AddEntryToMeal)
	Router.POST("/meal", models.AddMeal)

	// entry
	Router.GET("/entry/search", models.GetEntry)
	Router.POST("/entry/add", models.AddEntry)

}
