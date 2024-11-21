package router

import (
	"mauit/service"

	"github.com/gin-gonic/gin"
)

// Add endpoints to the API, noun [verb] object pathing
func SetRoutes(Router *gin.Engine) {

	// all dieters
	Router.GET("/dieters/all", service.GetDieters)
	Router.POST("/dieters", service.AddDieter)
	Router.DELETE("/dieters", service.DeleteDieter)

	// single dieter
	Router.GET("/dieter/name", service.GetDieter)
	Router.GET("/dieter/calories", service.GetDieterCalories)
	Router.GET("/dieter/remaining", service.GetRemainingDieterCalories)
    Router.GET("/dieter/meals", service.GetDieterMeals)
    Router.GET("/dieter/mealstoday", service.GetDieterMealsToday)
    Router.POST("/dieter/calories", service.SetDieterCalories)

	// meal
	Router.GET("/meal", service.GetMeal)
	Router.GET("/meal/calories", service.GetMealCalories)
	Router.GET("/meal/entries", service.GetMealEntries)
	Router.POST("/meal/entry", service.AddEntryToMeal)
	Router.POST("/meal", service.AddMeal)
	Router.DELETE("/meal", service.DeleteMeal)

	// entry
	Router.GET("/entry", service.GetEntry)
	Router.POST("/entry/add", service.AddEntry)
	Router.DELETE("/entry", service.DeleteEntry)

	// food
    Router.GET("/food/all", service.GetAllFood)
    Router.GET("/food", service.GetFood)
	Router.POST("/food", service.AddFood)
	Router.POST("/food/calories", service.EditFood)
	Router.DELETE("/food", service.DeleteFood)

}
