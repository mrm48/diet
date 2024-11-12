package models

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"mauit/mutils"
	"net/http"
	"time"
)

func GetDieters(req *gin.Context) {

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter")

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter rows from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create list of dieters from rows returned", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	defer rows.Close()

	req.IndentedJSON(http.StatusOK, Dieters)

	mutils.LogMessage("Request", "Dieters retrieved and sent to user")

}

// Add specifically a dieter
func AddDieter(req *gin.Context) {

	var dieter Dieter
	var newID int64

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from dieter").Scan(&newID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query dieter count from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO dieter values ($1, $2, $3)", newID+1, dieter.Calories, dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot store new dieter", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, dieter)

	mutils.LogMessage("Request", "Dieter added")

}

// Get dieter by name
func GetDieter(req *gin.Context) {

	var dieter Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE name=$1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get dieter information", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a list of dieters from search", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	defer rows.Close()

	for _, v := range Dieters {
		if v.Name == dieter.Name {
			req.IndentedJSON(http.StatusOK, v)
			mutils.LogMessage("Request", "Dieter information sent back to user")
			return
		}
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

// Set the calories available for a dieter
func SetDieterCalories(req *gin.Context) {

	var dieter Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter calories object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "UPDATE dieter SET Calories = $1 WHERE Name = $2", dieter.Calories, dieter.Name)

	if rows != nil {
		req.IndentedJSON(http.StatusOK, dieter)
		mutils.LogMessage("Request", "Calories updated for dieter")
		return
	} else if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot set dieter calories", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

func GetDieterCalories(req *gin.Context) {

	var dieter Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE Name = $1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a dieter object from search", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	if len(Dieters) == 1 {
		req.IndentedJSON(http.StatusOK, Dieters[0].Calories)
		return
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find unique Dieter requested", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}
}

func GetRemainingDieterCalories(req *gin.Context) {

	var dieter Dieter

	if err := req.BindJSON(&dieter); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * from dieter WHERE Name = $1", dieter.Name)

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	Dieter, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if Dieter != nil {
		rows, err := db.Query(context.Background(), "Select SUM(Calories) from meal WHERE dieter_id=$1 AND day=$2", dieter.ID, time.DateOnly)
		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
			return
		} else {
			if rows.Next() == true {
				err = rows.Scan(&dieter.Calories)
				if err != nil {
					mutils.LogApplicationError("Request", "Cannot parse sum of calories for this dieter", err)
					return
				} else {
					req.IndentedJSON(http.StatusOK, Dieter[1].Calories-dieter.Calories)
					return
				}
			}
		}
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find remaining dieter calories requested", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}
}

func GetMeal(req *gin.Context) {

	var meal Meal

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

	rows, err := db.Query(context.Background(), "Select * FROM meal WHERE name=$1 AND dieter=$2 AND day=$3", meal.Name, meal.Dieter, meal.Day)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot query meal from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[Meal])

	if meals != nil {
		mutils.LogMessage("Request", "Responded with the meal requested")
		req.IndentedJSON(http.StatusOK, meals)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

func GetMealCalories(req *gin.Context) {

	var meal Meal

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

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[Meal])

	meal.Calories = meals[0].Calories

	if meals != nil {
		mutils.LogMessage("Request", "Responded with the meal calories requested")
		req.IndentedJSON(http.StatusOK, meal)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

func AddMeal(req *gin.Context) {
	var meal Meal
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

	dieter, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

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

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[Entry])

	if entries != nil {
		return entries[0].ID
	}

	return 0
}

func GetEntry(req *gin.Context) {

	var entry Entry

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

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[Entry])

	if entries != nil {
		mutils.LogMessage("Request", "Responded with the entry requested")
		req.IndentedJSON(http.StatusOK, entries)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)
}

func AddEntry(req *gin.Context) {

	var entry Entry
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

	var entry Entry
	var meal []Meal

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

	meal, err = pgx.CollectRows(meals, pgx.RowToStructByName[Meal])

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

	var food Food

	var count int64

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food object from JSON provided", err)
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from food").Scan(&count)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query food count from database", err)
	}

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot get food count from database row returned", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Query(context.Background(), "INSERT INTO food (id, calories, units, name) values ($1, $2, $3, $4)", count+1, food.Calories, food.Units, food.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot insert food into database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, food)
	mutils.LogMessage("Request", "Added food to the database")

}

func EditFood(req *gin.Context) {

	var food Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food calories object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "UPDATE food SET Calories = $1 WHERE Name = $2", food.Calories, food.Name)

	if rows != nil {
		req.IndentedJSON(http.StatusOK, food)
		mutils.LogMessage("Request", "Calories updated for food")
		return
	} else if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot set food calories", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusNotFound, nil)

}

func DeleteFood(req *gin.Context) {

	var food Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food object from JSON provided", err)
		req.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Query(context.Background(), "DELETE FROM food WHERE Name = $1", food.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot delete food from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)

}

func DeleteMeal(req *gin.Context) {

	var meal Meal

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

	rows, err := db.Query(context.Background(), "SELECT * FROM meal WHERE Name = $1 AND Dieter = $2 AND Day = $3", meal.Name, meal.Dieter, meal.Day)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot find meal in database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	dbMeal, err := pgx.CollectRows(rows, pgx.RowToStructByName[Meal])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create meal object from row in database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	if dbMeal == nil {
		mutils.LogApplicationError("Database Error", "Cannot find meal in database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	meal.ID = dbMeal[0].ID

	deleteEntriesByMeal(meal.ID, req)

	_, err = db.Query(context.Background(), "DELETE FROM meal WHERE ID=$1", meal.ID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot meal food from database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusOK, nil)

}

func deleteEntriesByMeal(mealID int64, req *gin.Context) {

	meal, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = meal.Query(context.Background(), "DELETE FROM entry WHERE MEAL_ID=$1", mealID)

	return
}

func deleteMealsByDieter(dieterID int64, req *gin.Context) {

    meal, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = meal.Query(context.Background(), "DELETE FROM meal WHERE DIETER=$1", dieterID)

	return
}

func GetAllFood(req *gin.Context) {
	var food []Food

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "SELECT * FROM food")

	if rows != nil {
		food, err = pgx.CollectRows(rows, pgx.RowToStructByName[Food])
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
