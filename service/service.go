package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mauit/models"
	"mauit/mutils"
	"mauit/repositories"
	"net/http"
)

func GetDieters(req *gin.Context) {

	Dieters, err := repositories.GetAllDieters()

	if err != nil {
		mutils.LogApplicationError("Application Error", "Could not return the list of dieters from the database", err)
		req.IndentedJSON(http.StatusInternalServerError, errors.New("could not return the list of dieters from the database"))
		return
	}

	req.IndentedJSON(http.StatusOK, Dieters)
	mutils.LogMessage("Request", "Dieters retrieved and sent to user")

}

// AddDieter Add specifically a dieter
func AddDieter(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter object from JSON provided"))
		return
	}

	err := repositories.AddNewDieter(dieter)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot add user to the database", err)
		req.IndentedJSON(http.StatusInternalServerError, errors.New("cannot add user to the database"))
		return
	}

	req.IndentedJSON(http.StatusCreated, dieter)

	mutils.LogMessage("Request", "Dieter added")

}

// GetDieter Get dieter by name
func GetDieter(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter object from JSON provided"))
		return
	}

	dieter, err := repositories.GetSingleDieter(dieter)

	if err != nil {
		req.IndentedJSON(http.StatusNotFound, errors.New("cannot retrieve dieter just created"))
		return
	}

	req.IndentedJSON(http.StatusOK, dieter)

}

// SetDieterCalories Set the calories available for a dieter
func SetDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter calories object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter calories object from JSON provided"))
		return
	}

	err := repositories.UpdateDieterCalories(dieter)

	if err != nil {
		req.IndentedJSON(http.StatusNotFound, errors.New("cannot update calories for dieter"))
	}

	req.IndentedJSON(http.StatusOK, dieter)

}

func GetDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter object from JSON provided"))
		return
	}

	Dieters, err := repositories.GetDieterCalories(dieter)

	if err == nil {
		req.IndentedJSON(http.StatusOK, Dieters[0].Calories)
		return
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find unique Dieter requested", nil)
		req.IndentedJSON(http.StatusNotFound, errors.New("cannot find unique Dieter requested"))
		return
	}
}

func GetDieterMealsToday(req *gin.Context) {

	var dieter models.Dieter

	day := models.GetCurrentDate()

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter object from JSON provided"))
		return
	}

	meals, err := repositories.GetDieterMealsToday(dieter, day)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, errors.New("cannot get list of meals for today"))
		return
	}

	req.IndentedJSON(http.StatusOK, meals)

}

func GetRemainingDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	day := models.GetCurrentDate()

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter from information provided"))
		return
	}

	calories, err := repositories.GetRemainingCaloriesToday(dieter, day)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot get remaining calories for dieter", err)
		req.IndentedJSON(http.StatusInternalServerError, errors.New("cannot get remaining calories for dieter"))
		return
	}

	dieter.Calories = calories

	req.IndentedJSON(http.StatusOK, dieter)

}

func GetMeal(req *gin.Context) {

	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create meal object from JSON provided"))
		return
	}

	meals, err := repositories.GetMeal(meal)

	if meals != nil && err == nil {
		mutils.LogMessage("Request", "Responded with the meal requested")
		req.IndentedJSON(http.StatusOK, meals)
		return
	}

	req.IndentedJSON(http.StatusNotFound, errors.New("control error, contact system administrator"))

}

func GetMealCalories(req *gin.Context) {

	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create meal object from JSON provided"))
		return
	}

	newCalories, err := repositories.GetMealCalories(meal)
	meal.Calories = newCalories

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot get calories from meal database", err)
		req.IndentedJSON(http.StatusInternalServerError, errors.New("cannot get calories from meal database"))
		return
	}

	mutils.LogMessage("Request", "Responded with the meal calories requested")
	req.IndentedJSON(http.StatusOK, meal)
}

func GetMealEntries(req *gin.Context) {

	var meal models.Meal
	var entries []models.Entry

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	entries, err := repositories.GetMealEntries(meal)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot populate list of entries from rows returned", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, entries)

}

func GetDieterMeals(req *gin.Context) {

	var dieter models.Dieter
	var meals []models.Meal

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	meals, err := repositories.GetDieterMeals(dieter)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot populate list of meals from rows returned", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, meals)

}

func AddMeal(req *gin.Context) {
	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	// add in any missing fields to meal object (don't need day, dieterid or calories)
	if meal.Day == "" {
		meal.Day = models.GetCurrentDate()
	}

	err := repositories.AddMeal(meal)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	req.IndentedJSON(http.StatusCreated, meal)

}

func GetEntry(req *gin.Context) {

	var entry models.Entry

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	entry, err := repositories.GetEntry(entry)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, entry)
}

func AddEntry(req *gin.Context) {

	var entry models.Entry

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	entry, err := repositories.AddEntry(entry)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot add entry into database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, entry)
	mutils.LogMessage("Request", "Added entry to the database")

}

func AddEntryToMeal(req *gin.Context) {

	var entry models.Entry

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err := repositories.AddEntryToMeal(entry)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot update meal by adding the entry", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, entry)

}

func AddFood(req *gin.Context) {

	var food models.Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food object from JSON provided", err)
	}

	err := repositories.AddFoodRow(food)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot insert food into database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, food)
	mutils.LogMessage("Request", "Added food to the database")

}

func GetFood(req *gin.Context) {

	var food models.Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	food, err := repositories.GetFoodRow(food)

	if err != nil {
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	if food.Name != "nil" {
		req.IndentedJSON(http.StatusOK, food)
	}

}

func EditFood(req *gin.Context) {

	var food models.Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food calories object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err := repositories.UpdateFood(food)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot set food calories", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	} else {
		req.IndentedJSON(http.StatusOK, food)
		mutils.LogMessage("Request", "Calories updated for food")
		return
	}

}

func DeleteFood(req *gin.Context) {

	var food models.Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err := repositories.DeleteFoodRow(food)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot delete food from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)

}

func DeleteMeal(req *gin.Context) {

	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err := repositories.DeleteMeal(meal)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot delete meal from database", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)

}

func deleteMealsForDieter(dieterID int64, req *gin.Context) {

	err := repositories.DeleteMealsForDieter(dieterID)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)
	return

}

func DeleteMealEntries(req *gin.Context) {
	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	meals, err := repositories.GetMeal(meal)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	err = deleteEntriesByMeal(meals[0].ID, req)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Could not remove meal entries from database", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)

}

func deleteEntriesByMeal(mealID int64, req *gin.Context) error {

	err := repositories.DeleteEntriesByMeal(mealID)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, err)
		return err
	}

	return nil

}

func GetAllFood(req *gin.Context) {

	food, err := repositories.GetAllFood()

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot get all food from database", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, food)

}

func DeleteDieter(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err := repositories.DeleteDieter(dieter)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

}

func DeleteEntry(req *gin.Context) {

	var entry models.Entry

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	err := repositories.DeleteEntry(entry)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete entry by ID", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)

}
