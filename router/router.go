package router

import (
	"mauit/service"

	"github.com/gin-gonic/gin"
)

// SetRoutes to add endpoints to the API, noun [verb] object pathing
func SetRoutes(Router *gin.Engine) {

	// all dieters
	Router.GET("/dieters/all", service.GetDieters)
	Router.POST("/dieters", service.AddDieter)
	Router.DELETE("/dieters", service.DeleteDieter)

	// single dieter
	Router.POST("/dieter/calories", service.GetDieterCalories)
	Router.POST("/dieter/remaining", service.GetRemainingDieterCalories)
	Router.POST("/dieter/meals", service.GetDieterMeals)
	Router.POST("/dieter/mealstoday", service.GetDieterMealsToday)
	Router.POST("/dieter/name", service.GetDieter)
	Router.PUT("/dieter/calories", service.SetDieterCalories)

	// meal
	Router.POST("/meal", service.GetMeal)
	Router.POST("/meal/calories", service.GetMealCalories)
	Router.POST("/meal/entries", service.GetMealEntries)
	Router.POST("/meal/entry", service.AddEntryToMeal)
	Router.PUT("/meal", service.AddMeal)
	Router.PUT("/meal/calories", service.SetMealCalories)
	Router.DELETE("/meal", service.DeleteMeal)
	Router.DELETE("/meal/entries", service.DeleteMealEntries)

	// entry
	Router.POST("/entry", service.GetEntry)
	Router.PUT("/entry", service.AddEntry)
	Router.DELETE("/entry", service.DeleteEntry)

	// food
	Router.GET("/food/all", service.GetAllFood)
	Router.POST("/food", service.GetFood)
	Router.PUT("/food", service.AddFood)
	Router.PUT("/food/calories", service.EditFood)
	Router.DELETE("/food", service.DeleteFood)

}
