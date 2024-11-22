package models

import (
	"context"

	"mauit/mutils"

	"github.com/jackc/pgx/v5"
)

func getConnection() (*pgx.Conn, error) {
    db, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/meal")

    if err != nil {
        mutils.LogConnectionError(err)
        return nil, err
    }

    return db, err
}

func GetAllDieters() ([]Dieter, error) {
    
    db, err := getConnection()
	rows, err := db.Query(context.Background(), "Select * FROM dieter")

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter rows from database", err)
		return nil, err
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create list of dieters from rows returned", err)
		return nil, err
	}

	defer rows.Close()
    return Dieters, err

}

func GetSingleDieter(dieter Dieter) (Dieter, error) {

	db, err := getConnection()

	if err != nil {
		return dieter, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE name=$1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get dieter information", err)
		return dieter, err
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a list of dieters from search", err)
		return dieter, err
	}

	defer rows.Close()

	for _, v := range Dieters {
		if v.Name == dieter.Name {
			return v, nil
		}
	}

    return dieter, err
}

func AddNewDieter(dieter Dieter) (error) {

	var newID int64

	db, err := getConnection()

	if err != nil {
		mutils.LogConnectionError(err)
		return err
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from dieter").Scan(&newID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query dieter count from database", err)
		return err
	}

	_, err = db.Exec(context.Background(), "INSERT INTO dieter values ($1, $2, $3)", newID+1, dieter.Calories, dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot store new dieter", err)
		return err
	}

    return nil

}

func GetFoodRow(food Food) (Food, error) {

    var errorFood Food
    errorFood.Name = "nil"

	db, err := getConnection()

	rows, err := db.Query(context.Background(), "Select * FROM Food WHERE name=$1", food.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get food information", err)
		return errorFood, err
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[Food])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a list of food from search", err)
		return errorFood, err
	}

	defer rows.Close()

	for _, v := range items {
		if v.Name == food.Name {
			mutils.LogMessage("Request", "food information sent back to user")
			return v, err
		}
	}

    return errorFood, err
}

func AddFoodRow(food Food) error {

	db, err := getConnection()

    if err != nil {
        return err
    }

	var count int64

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from food").Scan(&count)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query food count from database", err)
        return err
	}

	_, err = db.Query(context.Background(), "INSERT INTO food (id, calories, units, name) values ($1, $2, $3, $4)", count+1, food.Calories, food.Units, food.Name)
    return err
}

func UpdateFood(food Food) error {

    db, err := getConnection()

    if err != nil {
        return err
    }

	_, err = db.Query(context.Background(), "UPDATE food SET Calories = $1 WHERE Name = $2", food.Calories, food.Name)

    return err 
}

func DeleteFoodRow(food Food) error {

    db, err  := getConnection()

	if err != nil {
        return err
	}

	_, err = db.Query(context.Background(), "DELETE FROM food WHERE Name = $1", food.Name)

    return err
}
