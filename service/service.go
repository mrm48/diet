package service

import (
	"context"
	"mauit/models"
	"mauit/mutils"
	"mauit/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetDieters(req *gin.Context) {

    Dieters, err := repositories.GetAllDieters()

    if err != nil {
        mutils.LogApplicationError("Application Error", "Could not return the list of dieters from the database", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

	req.IndentedJSON(http.StatusOK, Dieters)
	mutils.LogMessage("Request", "Dieters retrieved and sent to user")

}

// Add specifically a dieter
func AddDieter(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

    err := repositories.AddNewDieter(dieter)

    if err != nil {
        mutils.LogApplicationError("Application Error", "Cannot add user to the database", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }   

	req.IndentedJSON(http.StatusCreated, dieter)

	mutils.LogMessage("Request", "Dieter added")

}

// Get dieter by name
func GetDieter(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

    dieter, err := repositories.GetSingleDieter(dieter)

    if err != nil {
        req.IndentedJSON(http.StatusNotFound, nil)
        return
    }

    req.IndentedJSON(http.StatusOK, dieter)

}

// Set the calories available for a dieter
func SetDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter calories object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

    err := repositories.UpdateDieterCalories(dieter)

    if err != nil {
        req.IndentedJSON(http.StatusNotFound, nil)
    }

	req.IndentedJSON(http.StatusOK, dieter)

}

func GetDieterCalories(req *gin.Context) {

	var dieter models.Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

    Dieters, err := repositories.GetDieterCalories(dieter)

	if err == nil {
		req.IndentedJSON(http.StatusOK, Dieters[0].Calories)
		return
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find unique Dieter requested", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}
}

func GetDieterMealsToday(req *gin.Context) {

    var dieter models.Dieter

    day := models.GetCurrentDate()
    mutils.LogMessage("debug", day)

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

    meals, err := repositories.GetDieterMealsToday(dieter, day)

    if err != nil {
        req.IndentedJSON(http.StatusInternalServerError, err)
        return
    } 

    req.IndentedJSON(http.StatusOK, meals)

}

func GetRemainingDieterCalories(req *gin.Context) {

    var dieter models.Dieter

    day := models.GetCurrentDate()

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

}

func GetMeal(req *gin.Context) {

	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

    meals, err := repositories.GetMeal(meal)

	if meals != nil && err == nil {
		mutils.LogMessage("Request", "Responded with the meal requested")
		req.IndentedJSON(http.StatusOK, meals)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

func GetMealCalories(req *gin.Context) {

	var meal models.Meal

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select SUM(Calories) from meal WHERE name=$1 AND day=$2 AND dieter=3", meal.Name, meal.Day, meal.Dieter)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot query meal from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

	meal.Calories = meals[0].Calories

	if meals != nil {
		mutils.LogMessage("Request", "Responded with the meal calories requested")
		req.IndentedJSON(http.StatusOK, meal)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

func GetMealEntries(req *gin.Context) {

    var meal models.Meal
    var entries []models.Entry

    if err := req.BindJSON(&meal); err != nil {
        mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

    db, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/meal")

    if err != nil {
        mutils.LogConnectionError(err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

    rows, err := db.Query(context.Background(), "Select * from entry where MEAL_ID = $1", strconv.FormatInt(meal.ID, 10))

    if err != nil {
        mutils.LogApplicationError("Application Error", "Cannot find entries for provided meal ID", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

	entries, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Entry])

    if err != nil {
        mutils.LogApplicationError("Application Error", "Cannot populate list of entries from rows returned", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

    req.IndentedJSON(http.StatusOK, entries)
    return
}

func GetDieterMeals(req *gin.Context) {

    var dieter models.Dieter
    var meals []models.Meal

    if err := req.BindJSON(&dieter); err != nil {
        mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

    db, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/meal")

    if err != nil {
        mutils.LogConnectionError(err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

    rows, err := db.Query(context.Background(), "Select * from meal where dieter = $1", dieter.Name)

    if err != nil {
        mutils.LogApplicationError("Application Error", "Cannot find meals for provided dieter name", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

	meals, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

    if err != nil {
        mutils.LogApplicationError("Application Error", "Cannot populate list of meals from rows returned", err)
        req.IndentedJSON(http.StatusInternalServerError, nil)
        return
    }

    req.IndentedJSON(http.StatusOK, meals)
    return
}

func AddMeal(req *gin.Context) {
	var meal models.Meal
	var newID int64

	if err := req.BindJSON(&meal); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	// add in any missing fields to meal object (don't need day, dieterid or calories)
	if meal.Day == "" {
		meal.Day = time.DateOnly
	}

	if meal.Dieterid == 0 {
		meal.Dieterid = getDieterIDByName(meal.Dieter)
	}

	meal.Calories = 0

	if meal.Dieterid != 0 {
		db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

		if err != nil {
			mutils.LogConnectionError(err)
			req.IndentedJSON(http.StatusInternalServerError, nil)
			return
		}

		err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from meal").Scan(&newID)

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot query meal count from database", err)
			req.IndentedJSON(http.StatusInternalServerError, nil)
			return
		}

		_, err = db.Exec(context.Background(), "INSERT INTO meal values ($1, $2, $3, $4, $5, $6)", newID+1, meal.Calories, meal.Day, meal.Dieter, meal.Dieterid, meal.Name)

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot store new meal", err)
			req.IndentedJSON(http.StatusInternalServerError, nil)
			return
        }

        req.IndentedJSON(http.StatusCreated, meal)

		mutils.LogMessage("Request", "Meal added")
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find dieter id", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}
}

func getDieterIDByName(name string) int64 {
	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")
	if err != nil {
		mutils.LogConnectionError(err)
		return 0
	}
	rows, err := db.Query(context.Background(), "SELECT * FROM dieter WHERE NAME=$1", name)
	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query dieter from database", err)
		return 0
	}

	dieter, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if dieter != nil {
		return dieter[0].ID
	}

	return 0
}

func getMealCalories(id int64) int64 {
	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")
	if err != nil {
		mutils.LogConnectionError(err)
		return 0
	}
	rows, err := db.Query(context.Background(), "SUM(Calories) FROM entry WHERE MEAL_ID=$1", id)
	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query entries from database", err)
		return 0
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Entry])

	if entries != nil {
		return entries[0].ID
	}

	return 0
}

func GetEntry(req *gin.Context) {

	var entry models.Entry

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM entry WHERE ID=$1", entry.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot query entry from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Entry])

	if entries != nil {
		mutils.LogMessage("Request", "Responded with the entry requested")
		req.IndentedJSON(http.StatusOK, entries)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)
}

func AddEntry(req *gin.Context) {

	var entry models.Entry
	var newID int64

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from entry").Scan(&newID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query entry count from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Query(context.Background(), "INSERT INTO entry values ($1, $2, $3, $4)", newID+1, entry.Calories, entry.FoodID, entry.MealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot insert entry into database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, entry)
	mutils.LogMessage("Request", "Added entry to the database")

}

func AddEntryToMeal(req *gin.Context) {

	var entry models.Entry
	var meal []models.Meal

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	meals, err := db.Query(context.Background(), "SELECT * FROM meal WHERE ID = $1", entry.MealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query meal from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	meal, err = pgx.CollectRows(meals, pgx.RowToStructByName[models.Meal])

	if len(meal) != 1 {
		mutils.LogApplicationError("Application Error", "One and only one meal must match by meal ID", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	newCalories := entry.Calories + meal[0].Calories

	_, err = db.Query(context.Background(), "UPDATE meal SET Calories = $1 WHERE Meal_ID = $2", newCalories, entry.MealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot update meal in database", err)
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
	var food []models.Food

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	rows, err := db.Query(context.Background(), "SELECT * FROM food")

	if rows != nil {
		food, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Food])
		if err != nil {
			mutils.LogApplicationError("Application Error", "Cannot make a list of food from rows returned from database", err)
			req.IndentedJSON(http.StatusInternalServerError, nil)
			return
		}
		req.IndentedJSON(http.StatusOK, food)
		mutils.LogMessage("Request", "All food items returned")
		return
	} else if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get all food", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	} else {
		mutils.LogApplicationError("Application Error", "All food items returned, but the list is empty", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
	}

}

func DeleteDieter(req *gin.Context) {

	var dieter models.Dieter

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		mutils.LogApplicationError("Application Error", "Cannot connect to database to delete dieter", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err = db.QueryRow(context.Background(), "SELECT * from dieter WHERE Name=$1", dieter.Name).Scan(&dieter.ID, &dieter.Calories, &dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot retrieve dieter with name provided", err)
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	deleteMealsForDieter(dieter.ID, req)

	_, err = db.Query(context.Background(), "DELETE from dieter WHERE ID=$1", dieter.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete dieter retrieved by ID", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	return
}

func DeleteEntry(req *gin.Context) {

	var entry models.Entry

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	if err := req.BindJSON(&entry); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create entry object from JSON provided", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Query(context.Background(), "DELETE from ENTRY where ID = $1", entry.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete entry by ID", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)
	return
}
