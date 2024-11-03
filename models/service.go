package models

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"mauit/mutils"
	"net/http"
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

	Dieters = append(Dieters, dieter)

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
	Dieters, err = pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a dieter object from search", err)
		req.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	if len(Dieters) != 0 {
		req.IndentedJSON(http.StatusOK, Dieters[0].Calories)
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find Dieter requested", nil)
		req.IndentedJSON(http.StatusNotFound, nil)
	}
}

func GetMeal(req *gin.Context){
    req.IndentedJSON(http.StatusServiceUnavailable, nil)
    return
}

func GetEntry(req *gin.Context){
    req.IndentedJSON(http.StatusServiceUnavailable, nil)
    return
}

func AddEntry(req *gin.Context){
    req.IndentedJSON(http.StatusServiceUnavailable, nil)
    return
}

func AddEntryToMeal(req *gin.Context){
    req.IndentedJSON(http.StatusServiceUnavailable, nil)
    return
}
