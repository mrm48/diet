package repositories

import (
	"context"
	"errors"
	"strconv"

	"mauit/models"
	"mauit/mutils"

	"github.com/jackc/pgx/v5"
)

func getConnection() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/meal")

	if err != nil {
		mutils.LogConnectionError(err)
		return nil, errors.New("error 001: Cannot connect to the database")
	}

	return db, nil
}

func GetAllDieters() ([]models.Dieter, error) {

	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter")

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter rows from database", err)
		return nil, errors.New("error 101: Cannot get dieter information")
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create list of dieters from rows returned", err)
		return nil, errors.New("error 201: Cannot get dieter information")
	}

	defer rows.Close()
	return Dieters, err

}

func GetSingleDieter(dieter models.Dieter) (models.Dieter, error) {

	rows, err := retrieveDieter(dieter)

	if err == nil {
		Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

		if err != nil {
			mutils.LogApplicationError("Application Error", "Cannot create a list of dieters from search", err)
			return dieter, errors.New("error 201: Cannot create a list of dieters from search")
		}

		defer rows.Close()

		for _, v := range Dieters {
			if v.Name == dieter.Name {
				return v, nil
			}
		}
	}

	return dieter, err
}

func AddNewDieter(dieter models.Dieter) error {

	var newID int64

	db, err := getConnection()

	if err != nil {
		return err
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from dieter").Scan(&newID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query dieter count from database", err)
		return errors.New("error 101: Cannot get dieter count")
	}

	_, err = db.Exec(context.Background(), "INSERT INTO dieter values ($1, $2, $3)", newID+1, dieter.Calories, dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot store new dieter", err)
		return errors.New("error 102: Cannot store new dieter")
	}

	return nil

}

func UpdateDieterCalories(dieter models.Dieter) error {

	db, err := getConnection()

	if err != nil {
		return err
	}

	rows, err := db.Query(context.Background(), "UPDATE dieter SET Calories = $1 WHERE Name = $2", dieter.Calories, dieter.Name)

	if rows != nil {
		mutils.LogMessage("Request", "Calories updated for dieter")
		return nil
	} else if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot set dieter calories", err)
		return errors.New("error 102: Cannot update dieter information")
	}

	return nil
}

func GetDieterCalories(dieter models.Dieter) ([]models.Dieter, error) {

	rows, err := retrieveDieter(dieter)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
		return nil, errors.New("cannot get dieter information")
	}
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a dieter object from search", err)
		return nil, errors.New("cannot create a dieter object to return")
	}

	return Dieters, nil
}

func GetDieterMealsToday(dieter models.Dieter, day string) ([]models.Meal, error) {
	db, err := getConnection()

	if err != nil {
		return nil, errors.New("could not connect to the database")
	}

	rows, err := db.Query(context.Background(), "SELECT * from meal WHERE dieter=$1 AND day=$2", dieter.Name, day)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve meals by day for dieter from database", err)
		return nil, errors.New("cannot retrieve meals from database")
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot populate list of meals with data returned from database", err)
		return nil, errors.New("cannot create an array of objects to return")
	}

	return meals, nil
}

func GetRemainingCaloriesToday(dieter models.Dieter, day string) (int, error) {

	rows, err := retrieveDieter(dieter)

	if err != nil {
		mutils.LogApplicationError("Database error", "Cannot find dieter by name", err)
		return 0, errors.New("cannot get dieter information")
	}

	Dieter, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if Dieter != nil {

		dieter.ID = Dieter[0].ID

		db, err := getConnection()
		rows, err := db.Query(context.Background(), "SELECT * from meal WHERE dieterid=$1 AND day=$2", dieter.ID, day)

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot get meals from database for current user on the current day", err)
			return 0, errors.New("cannot get meals for the dieter")
		}

		meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

		if len(meals) > 0 {

			rows, err = db.Query(context.Background(), "Select SUM(Calories) from meal WHERE dieterid=$1 AND day=$2", dieter.ID, day)
			if err != nil {
				mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
				return 0, errors.New("cannot get today's calories for dieter")
			} else {
				if rows.Next() == true {
					err = rows.Scan(&dieter.Calories)
					if err != nil {
						mutils.LogApplicationError("Request", "Cannot parse sum of calories for this dieter", err)
						return 0, errors.New("cannot parse the sum of calories for this dieter")
					} else {
						return Dieter[0].Calories - dieter.Calories, nil
					}
				}
			}
		} else {
			return Dieter[0].Calories, nil
		}
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find remaining dieter calories requested", nil)
		return 0, errors.New("cannot find dieter information")
	}

	return 0, errors.New("error in control flow")
}

func DeleteDieter(dieter models.Dieter) error {
	db, err := getConnection()

	if err != nil {
		return errors.New("database connection error")
	}

	err = db.QueryRow(context.Background(), "SELECT * from dieter WHERE Name=$1", dieter.Name).Scan(&dieter.ID, &dieter.Calories, &dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot retrieve dieter with name provided", err)
		return errors.New("cannot find dieter")
	}

	err = DeleteMealsForDieter(dieter.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete meals with dieter id provided", err)
		return errors.New("cannot delete meals with dieter id provided")
	}

	_, err = db.Query(context.Background(), "DELETE from dieter WHERE ID=$1", dieter.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete dieter retrieved by ID", err)
		return errors.New("cannot delete dieter")
	}

	return nil
}

func DeleteMealsForDieter(dieterID int64) error {

	meal, err := getConnection()

	if err != nil {
		return errors.New("cannot open connection to the database")
	}

	rows, err := meal.Query(context.Background(), "SELECT ID FROM meal WHERE dieterID=$1", dieterID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot find meals by dieter from database", err)
		return errors.New("cannot find meals for requested dieter")
	}

	defer rows.Close()

	var index int64

	for rows.Next() {
		err = rows.Scan(&index)
		if err != nil {
			mutils.LogApplicationError("Application Error", "Cannot get meal ID from returned rows", err)
			return errors.New("cannot retrieve the meal ID: " + strconv.FormatInt(index, 10))
		}
		err = DeleteEntriesByMeal(index)

		if err != nil {
			mutils.LogApplicationError("Application Error", "Cannot delete entries for meal", err)
			return errors.New("cannot remove entries from the meal")
		}

		conn, err := getConnection()
		if err != nil {
			return err
		}
		_, err = conn.Query(context.Background(), "DELETE FROM meal WHERE ID=$1", index)
		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot delete meal from database", err)
			return errors.New("cannot remove meal from the database")
		}
	}

	return nil

}

func GetFoodRow(food models.Food) (models.Food, error) {

	var errorFood models.Food
	errorFood.Name = "nil"

	db, err := getConnection()

	if err != nil {
		return errorFood, errors.New("cannot connect to the database")
	}

	rows, err := db.Query(context.Background(), "Select * FROM Food WHERE name=$1", food.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get food information", err)
		return errorFood, errors.New("cannot retrieve food information from provided name:" + food.Name)
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Food])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a list of food from search", err)
		return errorFood, errors.New("cannot retrieve a list of food items matching the provided name: " + food.Name)
	}

	defer rows.Close()

	for _, v := range items {
		if v.Name == food.Name {
			mutils.LogMessage("Request", "food information sent back to user")
			return v, nil
		}
	}

	return errorFood, errors.New("broken control flow, should not be able to get here")
}

func GetMeal(meal models.Meal) ([]models.Meal, error) {
	db, err := getConnection()

	if err != nil {
		return nil, errors.New("cannot connect to the database")
	}

	rows, err := db.Query(context.Background(), "Select * FROM meal WHERE name=$1 AND dieter=$2 AND day=$3", meal.Name, meal.Dieter, meal.Day)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot query meal from database", err)
		return nil, errors.New("cannot retrieve meal from the database")
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

	return meals, nil
}

func DeleteEntriesByMeal(mealID int64) error {

	meal, err := getConnection()

	if err != nil {
		return errors.New("cannot connect to the database")
	}

	_, err = meal.Query(context.Background(), "DELETE FROM entry WHERE MEAL_ID=$1", mealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot delete entries for meal from database", err)
		return errors.New("cannot remove entries from the database for the provided meal id")
	}

	return nil

}

func DeleteMeal(meal models.Meal) error {
	db, err := getConnection()

	if err != nil {
		return errors.New("cannot connect to the database")
	}

	var dbMeal models.Meal
	err = db.QueryRow(context.Background(), "SELECT ID FROM meal WHERE Name = $1 AND Dieter = $2 AND Day = $3", meal.Name, meal.Dieter, meal.Day).Scan(&dbMeal.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot find meal in database", err)
		return errors.New("cannot find the requested meal in the database")
	}

	meal.ID = dbMeal.ID

	err = DeleteEntriesByMeal(meal.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete meal entries from database", err)
		return errors.New("cannot remove meal entries from the database")
	}

	_, err = db.Query(context.Background(), "DELETE FROM meal WHERE ID=$1", meal.ID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot delete meal from database", err)
		return errors.New("cannot remove meal " + strconv.FormatInt(meal.ID, 10) + " from the database")
	}

	return nil
}

func GetMealCalories(meal models.Meal) (int, error) {

	db, err := getConnection()

	if err != nil {
		return 0, err
	}

	rows, err := db.Query(context.Background(), "Select SUM(Calories) from meal WHERE name=$1 AND day=$2 AND dieter=3", meal.Name, meal.Day, meal.Dieter)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query meal from database", err)
		return 0, err
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot parse meal objects from row returned from the database", err)
		return 0, err
	}

	return meals[0].Calories, err

}

func GetMealEntries(meal models.Meal) ([]models.Entry, error) {

	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * from entry where MEAL_ID = $1", strconv.FormatInt(meal.ID, 10))

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot find entries for provided meal ID", err)
		return nil, err
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Entry])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot populate list of entries from rows returned", err)
		return nil, err
	}

	return entries, nil
}

func GetDieterMeals(dieter models.Dieter) ([]models.Meal, error) {

	db, err := getConnection()

	if err != nil {
		mutils.LogConnectionError(err)
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * from meal where dieter = $1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot find meals for provided dieter name", err)
		return nil, err
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot populate list of meals from rows returned", err)
		return nil, err
	}

	return meals, nil

}

func AddMeal(meal models.Meal) error {
	var newID int64
	var dieter models.Dieter

	dieter.Name = meal.Dieter

	dieter, err := GetSingleDieter(dieter)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot find dieter in the database", err)
		return errors.New("cannot find dieter in the database")
	}

	meal.Calories = 0

	if meal.Dieterid != 0 {
		db, err := getConnection()

		if err != nil {
			mutils.LogConnectionError(err)
			return err
		}

		err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from meal").Scan(&newID)

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot query meal count from database", err)
			return err
		}

		_, err = db.Exec(context.Background(), "INSERT INTO meal values ($1, $2, $3, $4, $5, $6)", newID+1, meal.Calories, meal.Day, meal.Dieter, meal.Dieterid, meal.Name)

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot store new meal", err)
			return err
		}

		mutils.LogMessage("Request", "Meal added")
		return nil
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find dieter id", nil)
		return errors.New("cannot find dieter id")
	}
}

func getMealCalories(id int64) int64 {
	db, err := getConnection()
	if err != nil {
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

func getDieterIDByName(name string) int64 {
	db, err := getConnection()
	if err != nil {
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

func GetAllFood() ([]models.Food, error) {

	var food []models.Food

	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "SELECT * FROM food")

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query all food items from database", err)
		return nil, err
	}

	food, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Food])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot make a list of food from rows returned from database", err)
		return nil, err
	}
	return food, nil

}

func AddFoodRow(food models.Food) error {

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

func UpdateFood(food models.Food) error {

	db, err := getConnection()

	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "UPDATE food SET Calories = $1 WHERE Name = $2", food.Calories, food.Name)

	return err
}

func DeleteFoodRow(food models.Food) error {

	db, err := getConnection()

	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "DELETE FROM food WHERE Name = $1", food.Name)

	return err
}

func AddEntry(entry models.Entry) (models.Entry, error) {

	var newID int64
	var newEntry models.Entry

	db, err := getConnection()

	if err != nil {
		mutils.LogConnectionError(err)
		return newEntry, err
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from entry").Scan(&newID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query entry count from database", err)
		return newEntry, err
	}

	_, err = db.Query(context.Background(), "INSERT INTO entry values ($1, $2, $3, $4)", newID+1, entry.Calories, entry.FoodID, entry.MealID)

	return entry, nil

}

func AddEntryToMeal(entry models.Entry) error {

	var meal []models.Meal

	db, err := getConnection()

	if err != nil {
		return err
	}

	meals, err := db.Query(context.Background(), "SELECT * FROM meal WHERE ID = $1", entry.MealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query meal from database", err)
		return err
	}

	meal, err = pgx.CollectRows(meals, pgx.RowToStructByName[models.Meal])

	newCalories := entry.Calories + meal[0].Calories

	_, err = db.Query(context.Background(), "UPDATE meal SET Calories = $1 WHERE Meal_ID = $2", newCalories, entry.MealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot update meal in database", err)
		return err
	}

	return nil
}

func GetEntry(entry models.Entry) (models.Entry, error) {

	db, err := getConnection()

	if err != nil {
		return entry, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM entry WHERE ID=$1", entry.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot query entry from database", err)
		return entry, err
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Entry])

	if err != nil {
		mutils.LogMessage("Request", "Responded with the entry requested")
		return entry, err
	}

	return entries[0], nil

}

func DeleteEntry(entry models.Entry) error {

	db, err := getConnection()

	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "DELETE from ENTRY where ID = $1", entry.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete entry by ID", err)
		return err
	}

	return nil
}

func retrieveDieter(dieter models.Dieter) (pgx.Rows, error) {

	db, err := getConnection()

	if err != nil {
		mutils.LogConnectionError(err)
		return nil, errors.New("cannot connect to database")
	}

	rows, err := db.Query(context.Background(), "SELECT * FROM dieter WHERE name = $1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot query dieter from database", err)
		return nil, errors.New("cannot query dieter from database")
	}

	return rows, nil

}
