package models

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

func GetDieters(ctxt *gin.Context) {

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		LogConnectionError(err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter")

	if err != nil {
		LogApplicationError("Database Error", "Cannot retrieve dieter rows from database", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		LogApplicationError("Application Error", "Cannot create list of dieters from rows returned", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	defer rows.Close()

	ctxt.IndentedJSON(http.StatusOK, Dieters)

	LogMessage("Dieters retrieved and sent to user")

}

// Add specifically a dieter
func AddDieter(ctxt *gin.Context) {

	var dieter Dieter

	if err := ctxt.BindJSON(&dieter); err != nil {
		LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		ctxt.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	Dieters = append(Dieters, dieter)

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		LogConnectionError(err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO dieter values ($1, $2, $3)", dieter.ID, dieter.Calories, dieter.Name)

	if err != nil {
		LogApplicationError("Database Error", "Cannot store new dieter", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	ctxt.IndentedJSON(http.StatusCreated, dieter)

}

// Get dieter by name
func GetDieter(ctxt *gin.Context) {

	var dieter Dieter

	if err := ctxt.BindJSON(&dieter); err != nil {
		LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		ctxt.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		LogConnectionError(err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE name=$1", dieter.Name)

	if err != nil {
		LogApplicationError("Database Error", "Cannot get dieter information", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		LogApplicationError("Application Error", "Cannot create a list of dieters from search", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	defer rows.Close()

	for _, v := range Dieters {
		if v.Name == dieter.Name {
			ctxt.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	ctxt.IndentedJSON(http.StatusNotFound, nil)

}

// Set the calories available for a dieter
func SetDieterCalories(ctxt *gin.Context) {

	var dieter Dieter

	if err := ctxt.BindJSON(&dieter); err != nil {
		LogApplicationError("Application Error", "Cannot create dieter calories object from JSON provided", err)
		ctxt.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		LogConnectionError(err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "UPDATE dieter SET Calories = $1 WHERE Name = $2", dieter.Calories, dieter.Name)

	if rows != nil {
		SetCalories(dieter, dieter.Calories)
		ctxt.IndentedJSON(http.StatusOK, dieter)
		return
	} else if err != nil {
		LogApplicationError("Database Error", "Cannot set dieter calories", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	ctxt.IndentedJSON(http.StatusNotFound, nil)

}

func GetDieterCalories(ctxt *gin.Context) {

	var dieter Dieter

	if err := ctxt.BindJSON(&dieter); err != nil {
		LogApplicationError("Application Error", "Cannot create dieter object from JSON provided", err)
		ctxt.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		LogConnectionError(err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE Name = $1", dieter.Name)

	if err != nil {
		LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	Dieters, err = pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		LogApplicationError("Application Error", "Cannot create a dieter object from search", err)
		ctxt.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	if len(Dieters) != 0 {
		ctxt.IndentedJSON(http.StatusOK, Dieters[0].Calories)
	} else {
		LogApplicationError("Database Error", "Cannot find Dieter requested", nil)
		ctxt.IndentedJSON(http.StatusNotFound, nil)
	}
}

// Set the saved dieter's number of maximum calories
func SetCalories(d Dieter, c int) {

	for k, v := range Entries {
		if v.ID == GetID(d) {
			Dieters[k].Calories = c
		}
	}

}

// Get the unique ID for the dieter by name
func GetID(d Dieter) int64 {

	for _, v := range Dieters {
		if v.Name == d.Name {
			return v.ID
		}
	}

	return 0

}

// Set the unique ID for a dieter by name
func SetID(d Dieter) {

	for k, v := range Dieters {
		if v.Name == d.Name {
			Dieters[k].ID = d.ID
		}
	}

}

func LogConnectionError(err error) {

	slog := fmt.Sprintf("Cannot connect to the database: %v", err)
	log.Output(1, slog)

}

func LogApplicationError(ltype string, message string, err error) {

	slog := fmt.Sprintf("%v: %v : %v", ltype, message, err)
	log.Output(2, slog)

}

func LogMessage(message string) {

	log.Output(1, message)

}
