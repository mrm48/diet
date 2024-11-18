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
	Router.DELETE("/dieters", models.DeleteDieter)

	// single dieter
	Router.GET("/dieter/name", models.GetDieter)
	Router.GET("/dieter/calories", models.GetDieterCalories)
	Router.GET("/dieter/remaining", models.GetRemainingDieterCalories)
    Router.GET("/dieter/meals", models.GetDieterMeals)
    Router.GET("/dieter/mealstoday", models.GetDieterMealsToday)
    Router.POST("/dieter/calories", models.SetDieterCalories)

	// meal
	Router.GET("/meal", models.GetMeal)
	Router.GET("/meal/calories", models.GetMealCalories)
	Router.GET("/meal/entries", models.GetMealEntries)
	Router.POST("/meal/entry", models.AddEntryToMeal)
	Router.POST("/meal", models.AddMeal)
	Router.DELETE("/meal", models.DeleteMeal)

	// entry
	Router.GET("/entry", models.GetEntry)
	Router.POST("/entry/add", models.AddEntry)
	Router.DELETE("/entry", models.DeleteEntry)

	// food
    Router.GET("/food/all", models.GetAllFood)
	Router.POST("/food", models.AddFood)
	Router.POST("/food/calories", models.EditFood)
	Router.DELETE("/food", models.DeleteFood)

}
