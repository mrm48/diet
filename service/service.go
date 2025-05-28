package service

import (
	"errors"
	"mauit/models"
	"mauit/mutils"
	"mauit/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDieters from the database. This will return all dieters in a JSON array. Nothing needs to be sent in.
func GetDieters(req *gin.Context) {

	Dieters, err := repositories.GetAllDieters()

	req, err = mutils.WrapServiceError(err, "could not return the list of dieters from the database", req, http.StatusInternalServerError)
	if err == nil {
		req.IndentedJSON(http.StatusOK, Dieters)
		mutils.LogMessage("Request", "Dieters retrieved and sent to user")
		return
	}

}

// AddDieter to the database. The request must have the name and number of calories the new dieter will target each day.
func AddDieter(req *gin.Context) {
	var dieter models.Dieter

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter object from JSON provided", req, http.StatusBadRequest)

	if err == nil {

		err = repositories.AddNewDieter(dieter)

		req, err = mutils.WrapServiceError(err, "cannot add user to the database", req, http.StatusInternalServerError)
		if err == nil {
			req.IndentedJSON(http.StatusCreated, dieter)
			mutils.LogMessage("Request", "Dieter added")
		}

	}

}

// GetDieter from the database using the dieter name. Two dieters cannot have the same name.
func GetDieter(req *gin.Context) {

	var dieter models.Dieter

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter object from JSON provided", req, http.StatusBadRequest)
	if err == nil {
		dieter, err = repositories.GetSingleDieter(dieter)
		req, err = mutils.WrapServiceError(err, "cannot find dieter in database", req, http.StatusNotFound)

		if err == nil {
			req.IndentedJSON(http.StatusOK, dieter)
		}
	}

}

// SetDieterCalories in the database. This will set the target number of calories for a user using its name.
func SetDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter calories object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.UpdateDieterCalories(dieter)
		req, err = mutils.WrapServiceError(err, "cannot update dieter calories object", req, http.StatusNotFound)

		if err == nil {
			req.IndentedJSON(http.StatusOK, dieter)
		}
	}

}

// GetDieterCalories from the database, this will get the target number of calories from the user by name.
func GetDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "Cannot create dieter object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		req.IndentedJSON(http.StatusBadRequest, errors.New("cannot create dieter object from JSON provided"))
		Dieters, err := repositories.GetDieterCalories(dieter)

		req, err = mutils.WrapServiceError(err, "cannot find unique Dieter requested", req, http.StatusNotFound)

		if err == nil {
			req.IndentedJSON(http.StatusOK, Dieters[0].Calories)
		}
	}

}

// GetDieterMealsToday will return the meals consumed today by a user by name.
func GetDieterMealsToday(req *gin.Context) {
	var dieter models.Dieter
	day := models.GetCurrentDate()

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		meals, err := repositories.GetDieterMealsToday(dieter, day)
		req, err = mutils.WrapServiceError(err, "cannot get list of meals for today", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, meals)
		}
	}
}

// GetRemainingDieterCalories from the database. This will get the number of calories remaining before the user (by name) gets their daily target.
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

// GetMeal from the database. Requires the meal name and day. If no day is provided, it defaults to the current day.
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

// GetMealCalories from the database for a single meal. Requires the meal name, dieter name and day.
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

// GetMealEntries for a single meal. This will return the entries consumed during the specified meal on the specified day for the specified user in a JSON array.
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

// GetDieterMeals by specifying the dieter. This will return the meals associated with the dieter in the request body
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

// AddMeal to the database using the meal object provided in the body of the request (name and dieter id).
func AddMeal(req *gin.Context) {
	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	// add in any missing fields to the meal object (don't need day, dieterid or calories)
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

// GetEntry from the database using the Entry object (ID) provided in the request.
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

// AddEntry to the database using (calories, food, meal ID) provided in the body of the request.
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

// AddEntryToMeal calories consumed using the (mealID, name and calories) provided in the body of the request
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

// AddFood item to the database. Requires the food "name", number of "calories" and the number of "units" corresponding to the number of calories.
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

// GetFood from the database by matching the name from the request body.
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

// EditFood calories by matching the food name specified in the request.
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

// DeleteFood from the database using the food object specified in the request body
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

// DeleteMeal from the database using the meal specified in the request
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

// deleteMealsForDieter specified by the request.
func deleteMealsForDieter(dieterID int64, req *gin.Context) {

	err := repositories.DeleteMealsForDieter(dieterID)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)
	return

}

// DeleteMealEntries associated with the meal object in the request.
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

// deleteEntriesByMeal ID given, helps with DeleteMealEntries.
func deleteEntriesByMeal(mealID int64, req *gin.Context) error {

	err := repositories.DeleteEntriesByMeal(mealID)

	if err != nil {
		req.IndentedJSON(http.StatusInternalServerError, err)
		return err
	}

	return nil

}

// GetAllFood items from the database and return as a JSON array
func GetAllFood(req *gin.Context) {

	food, err := repositories.GetAllFood()

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot get all food from database", err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	req.IndentedJSON(http.StatusOK, food)

}

// DeleteDieter from the database with the specified username.
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

// DeleteEntry from the database using Entry specified by the user (day, meal, dieter name)
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
