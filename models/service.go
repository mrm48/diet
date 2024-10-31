package models

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

func GetDieters(ctxt *gin.Context) {

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter")
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	ctxt.IndentedJSON(http.StatusOK, Dieters)

}

// Add specifically a dieter
func AddDieter(ctxt *gin.Context) {

	var dieter Dieter

	if err := ctxt.BindJSON(&dieter); err != nil {
		log.Fatal(err)
		return
	}

	Dieters = append(Dieters, dieter)

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		log.Fatal(err)
	}

	//query := "INSERT INTO dieter values (" + strconv.FormatInt(dieter.ID, 10) + "," + strconv.Itoa(dieter.Calories) + "," + "'" + dieter.Name + "'" + ")"

	_, err = db.Exec(context.Background(), "INSERT INTO dieter values ($1, $2, $3)", dieter.ID, dieter.Calories, dieter.Name)

	if err != nil {
		log.Fatal(err)
	}

	ctxt.IndentedJSON(http.StatusCreated, dieter)

}

// Get dieter by name
func GetDieter(ctxt *gin.Context) {

	var dieter Dieter

	if err := ctxt.BindJSON(&dieter); err != nil {
		return
	}

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE name=$1", dieter.Name)
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		log.Fatal(err)
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
		return
	}

	for _, v := range Dieters {
		if v.Name == dieter.Name {
			SetCalories(v, dieter.Calories)
			v.Calories = dieter.Calories
			ctxt.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	ctxt.IndentedJSON(http.StatusNotFound, nil)

}

// Set the saved dieter's number of maximum calories
func SetCalories(d Dieter, c int) {

	for k, v := range Entries {
		if v.ID == GetID(d) {
			Dieters[k].Calories = c
		}
	}

}

// Get maximum number of calories for a dieter
func GetCalories(d Dieter) int {

	for _, v := range Dieters {
		if v.ID == d.ID {
			return v.Calories
		}
	}

	return 0

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
