package router

import (
	"mauit/models"

	"github.com/gin-gonic/gin"
)

// Add endpoints to the API, noun object pathing
func SetRoutes(Router *gin.Engine) {

	// all dieters
	Router.GET("/dieters/all", models.GetDieters)
	Router.POST("/dieters", models.AddDieter)

	// single dieter
	Router.GET("/dieter/name", models.GetDieter)
	Router.POST("/dieter/calories", models.SetDieterCalories)
	Router.GET("/dieter/calories", models.GetDieterCalories)
	Router.GET("/dieter/remaining", models.GetRemainingDieterCalories)

	// meal
	Router.GET("/meal", models.GetMeal)
	Router.POST("/meal/entry", models.AddEntryToMeal)
	Router.POST("/meal", models.AddMeal)
	Router.GET("/meal/calories", models.GetMealCalories)

	// entry
	Router.GET("/entry/", models.GetEntry)
	Router.POST("/entry/add", models.AddEntry)

	// food
	Router.POST("/food/", models.AddFood)
	Router.POST("/food/calories", models.EditFood)
	Router.GET("/food/all", models.GetAllFood)
	Router.DELETE("/food", models.DeleteFood)

}
