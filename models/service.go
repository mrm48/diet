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

	_, err = db.Exec(context.Background(), "INSERT INTO dieter values ($1, $2, $3)", dieter.ID, dieter.Calories, dieter.Name)

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
			rows.Scan(&dieter.Calories)
			req.IndentedJSON(http.StatusOK, Dieter[1].Calories-dieter.Calories)
			return
		}
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find remaining dieter calories requested", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
		return
	}
	req.IndentedJSON(http.StatusNotFound, nil)
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

func AddMeal(req *gin.Context) {
	var meal Meal

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

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO meal values ($1, $2, $3, $4, $5)", meal.Calories, meal.Day, meal.Dieter, meal.Dieterid, meal.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot store new meal", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, meal)

	mutils.LogMessage("Request", "Meal added")

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

	_, err = db.Query(context.Background(), "INSERT INTO entry values ($1, $2, $3)", entry.Calories, entry.FoodID, entry.MealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot insert entry into database", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	req.IndentedJSON(http.StatusCreated, entry)
	mutils.LogMessage("Request", "Added entry to the database")

}

func AddEntryToMeal(req *gin.Context) {
	req.IndentedJSON(http.StatusServiceUnavailable, nil)
}

func AddFood(req *gin.Context) {

	var food Food

	if err := req.BindJSON(&food); err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create food object from JSON provided", err)
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Query(context.Background(), "INSERT INTO food values ($1, $2, $3)", food.Name, food.Calories, food.Units)

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
