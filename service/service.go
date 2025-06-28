package service

import (
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

			dieter, err = repositories.GetSingleDieter(dieter)

			req, err = mutils.WrapServiceError(err, "cannot retrieve user from database after it has been added", req, http.StatusInternalServerError)
			if err == nil {
				req.IndentedJSON(http.StatusCreated, dieter)
				mutils.LogMessage("Request", "Dieter added")
			}

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
	day := mutils.GetCurrentDate()

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

func GetRemainingDieterCalories(req *gin.Context) {
	var dieter models.Dieter
	day := mutils.GetCurrentDate()

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		calories, err := repositories.GetRemainingCaloriesToday(dieter, day)
		req, err = mutils.WrapServiceError(err, "cannot get remaining calories for dieter", req, http.StatusInternalServerError)

		if err == nil {
			dieter.Calories = calories
			req.IndentedJSON(http.StatusOK, dieter)
		}
	}
}

// GetMeal from the database. Requires the meal name and day. If no day is provided, it defaults to the current day.
func GetMeal(req *gin.Context) {
	var meal models.Meal

	err := req.BindJSON(&meal)
	req, err = mutils.WrapServiceError(err, "cannot create meal object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		meals, err := repositories.GetMeal(meal)
		req, err = mutils.WrapServiceError(err, "cannot retrieve meal from database", req, http.StatusNotFound)

		if err == nil {
			mutils.LogMessage("Request", "Responded with the meal requested")
			req.IndentedJSON(http.StatusOK, meals)
		}
	}
}

// GetMealCalories from the database for a single meal. Requires the meal name, dieter name and day.
func GetMealCalories(req *gin.Context) {
	var meal models.Meal

	err := req.BindJSON(&meal)
	req, err = mutils.WrapServiceError(err, "cannot create meal object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		newCalories, err := repositories.GetMealCalories(meal)
		req, err = mutils.WrapServiceError(err, "cannot get calories from meal database", req, http.StatusInternalServerError)

		if err == nil {
			meal.Calories = newCalories
			mutils.LogMessage("Request", "Responded with the meal calories requested")
			req.IndentedJSON(http.StatusOK, meal)
		}
	}
}

// GetMealEntries for a single meal. This will return the entries consumed during the specified meal on the specified day for the specified user in a JSON array.
func GetMealEntries(req *gin.Context) {
	var meal models.Meal

	err := req.BindJSON(&meal)
	req, err = mutils.WrapServiceError(err, "cannot create meal object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		entries, err := repositories.GetMealEntries(meal)
		req, err = mutils.WrapServiceError(err, "cannot populate list of entries from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, entries)
		}
	}
}

// GetDieterMeals by specifying the dieter. This will return the meals associated with the dieter in the request body
func GetDieterMeals(req *gin.Context) {
	var dieter models.Dieter

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		meals, err := repositories.GetDieterMeals(dieter)
		req, err = mutils.WrapServiceError(err, "cannot populate list of meals from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, meals)
		}
	}
}

func AddMeal(req *gin.Context) {
	var meal models.Meal

	err := req.BindJSON(&meal)
	req, err = mutils.WrapServiceError(err, "cannot create meal object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		if meal.Day == "" {
			meal.Day = mutils.GetCurrentDate()
		}

		err = repositories.AddMeal(meal)
		req, err = mutils.WrapServiceError(err, "cannot add meal to database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusCreated, meal)
		}
	}
}

func GetEntry(req *gin.Context) {
	var entry models.Entry

	err := req.BindJSON(&entry)
	req, err = mutils.WrapServiceError(err, "cannot create entry object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		entry, err = repositories.GetEntry(entry)
		req, err = mutils.WrapServiceError(err, "cannot retrieve entry from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, entry)
		}
	}
}

func AddEntry(req *gin.Context) {
	var entry models.Entry

	err := req.BindJSON(&entry)
	req, err = mutils.WrapServiceError(err, "cannot create entry object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		entry, err = repositories.AddEntry(entry)
		req, err = mutils.WrapServiceError(err, "cannot add entry to database", req, http.StatusInternalServerError)

		if err == nil {
			mutils.LogMessage("Request", "Added entry to the database")
			req.IndentedJSON(http.StatusCreated, entry)
		}
	}
}

func AddEntryToMeal(req *gin.Context) {
	var entry models.Entry

	err := req.BindJSON(&entry)
	req, err = mutils.WrapServiceError(err, "cannot create entry object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.AddEntryToMeal(entry)
		req, err = mutils.WrapServiceError(err, "cannot update meal by adding the entry", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, entry)
		}
	}
}

func AddFood(req *gin.Context) {
	var food models.Food

	err := req.BindJSON(&food)
	req, err = mutils.WrapServiceError(err, "cannot create food object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.AddFoodRow(food)
		req, err = mutils.WrapServiceError(err, "cannot insert food into database", req, http.StatusInternalServerError)

		if err == nil {
			mutils.LogMessage("Request", "Added food to the database")
			req.IndentedJSON(http.StatusCreated, food)
		}
	}
}

func GetFood(req *gin.Context) {
	var food models.Food

	err := req.BindJSON(&food)
	req, err = mutils.WrapServiceError(err, "cannot create food object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		food, err = repositories.GetFoodRow(food)
		req, err = mutils.WrapServiceError(err, "cannot retrieve food from database", req, http.StatusNotFound)

		if err == nil {
			if food.Name != "nil" {
				req.IndentedJSON(http.StatusOK, food)
			}
		}
	}
}

func EditFood(req *gin.Context) {
	var food models.Food

	err := req.BindJSON(&food)
	req, err = mutils.WrapServiceError(err, "cannot create food calories object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.UpdateFood(food)
		req, err = mutils.WrapServiceError(err, "cannot set food calories", req, http.StatusInternalServerError)

		if err == nil {
			mutils.LogMessage("Request", "Calories updated for food")
			req.IndentedJSON(http.StatusOK, food)
		}
	}
}

func DeleteFood(req *gin.Context) {
	var food models.Food

	err := req.BindJSON(&food)
	req, err = mutils.WrapServiceError(err, "cannot create food object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.DeleteFoodRow(food)
		req, err = mutils.WrapServiceError(err, "cannot delete food from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, nil)
		}
	}
}

// DeleteMeal from the database using the meal specified in the request
func DeleteMeal(req *gin.Context) {
	var meal models.Meal

	err := req.BindJSON(&meal)
	req, err = mutils.WrapServiceError(err, "cannot create meal object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.DeleteMeal(meal)
		req, err = mutils.WrapServiceError(err, "cannot delete meal from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, nil)
		}
	}
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

func DeleteMealEntries(req *gin.Context) {
	var meal models.Meal

	err := req.BindJSON(&meal)
	req, err = mutils.WrapServiceError(err, "cannot create meal object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		meals, err := repositories.GetMeal(meal)
		req, err = mutils.WrapServiceError(err, "cannot retrieve meal from database", req, http.StatusInternalServerError)

		if err == nil {
			err = deleteEntriesByMeal(meals[0].ID, req)
			req, err = mutils.WrapServiceError(err, "could not remove meal entries from database", req, http.StatusInternalServerError)

			if err == nil {
				req.IndentedJSON(http.StatusOK, nil)
			}
		}
	}
}

// deleteEntriesByMeal ID given, helps with DeleteMealEntries.
func deleteEntriesByMeal(mealID int64, req *gin.Context) error {
	err := repositories.DeleteEntriesByMeal(mealID)
	req, err = mutils.WrapServiceError(err, "cannot delete entries for meal", req, http.StatusInternalServerError)
	return err
}

// GetAllFood items from the database and return as a JSON array
func GetAllFood(req *gin.Context) {
	food, err := repositories.GetAllFood()
	req, err = mutils.WrapServiceError(err, "cannot get all food from database", req, http.StatusInternalServerError)

	if err == nil {
		req.IndentedJSON(http.StatusOK, food)
	}
}

func DeleteDieter(req *gin.Context) {
	var dieter models.Dieter

	err := req.BindJSON(&dieter)
	req, err = mutils.WrapServiceError(err, "cannot create dieter object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.DeleteDieter(dieter)
		req, err = mutils.WrapServiceError(err, "cannot delete dieter from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, nil)
		}
	}
}

func DeleteEntry(req *gin.Context) {
	var entry models.Entry

	err := req.BindJSON(&entry)
	req, err = mutils.WrapServiceError(err, "cannot create entry object from JSON provided", req, http.StatusBadRequest)

	if err == nil {
		err = repositories.DeleteEntry(entry)
		req, err = mutils.WrapServiceError(err, "cannot delete entry from database", req, http.StatusInternalServerError)

		if err == nil {
			req.IndentedJSON(http.StatusOK, nil)
		}
	}
}
