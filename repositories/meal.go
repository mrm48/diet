package repositories

import (
	"context"
	"errors"
	"strconv"

	"mauit/models"
	"mauit/mutils"

	"github.com/jackc/pgx/v5"
)

// getConnection establishes and returns a connection to the PostgreSQL database.
func getConnection() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/meal")
	return db, mutils.WrapError(err, "error 001: Cannot connect to the database", "connection")
}

// GetAllDieters retrieves all dieters from the database.
func GetAllDieters() ([]models.Dieter, error) {

	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter")

	err = mutils.WrapError(err, "error 101: cannot retrieve dieter rows from database", "table")
	if err == nil {
		Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

		err = mutils.WrapError(err, "error 201: Cannot get dieter information", "notfound")

		defer rows.Close()
		return Dieters, err
	}

	return nil, nil
}

// GetSingleDieter retrieves a single dieter from the database based on the provided dieter model.
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

// AddNewDieter adds a new dieter to the database.
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

// UpdateDieterCalories updates the calories for a dieter in the database.
func UpdateDieterCalories(dieter models.Dieter) error {

	db, err := getConnection()

	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "UPDATE dieter SET Calories = $1 WHERE Name = $2", dieter.Calories, dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot set dieter calories", err)
		return errors.New("error 102: Cannot update dieter information")
	}

	return nil
}

// GetDieterCalories retrieves the calories for a dieter from the database.
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

// GetDieterMealsToday retrieves all meals for a dieter on a specific day from the database.
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

// GetRemainingCaloriesToday gets the remaining calories for a 'dieter' on a specified 'day' from the database.
func GetRemainingCaloriesToday(dieter models.Dieter, day string) (int, error) {

	rows, err := retrieveDieter(dieter)

	if err != nil {
		mutils.LogApplicationError("Database error", "Cannot find dieter by name", err)
		return 0, errors.New("cannot get dieter information")
	}

	Dieter, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot parse dieter from row returned", err)
		return 0, errors.New("cannot parse dieter from row returned")
	}

	// Dieter is found, query all meals from dieter on specified day
	if Dieter != nil {

		dieter.ID = Dieter[0].ID

		db, err := getConnection()

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot connect to the database", err)
			return 0, errors.New("cannot connect to the database")
		}

		rows, err := db.Query(context.Background(), "SELECT * from meal WHERE dieterid=$1 AND day=$2", dieter.ID, day)

		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot get meals from database for current user on the current day", err)
			return 0, errors.New("cannot get meals for the dieter")
		}

		meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

		if err != nil {
			mutils.LogApplicationError("Application Error", "Cannot parse meals from row returned", err)
			return 0, errors.New("cannot parse meals from row returned")
		}

		// Found meals from today for this dieter, query the database for the sum of calories for those meals.
		// TODO: is the above query to check if there are meals for the "day" necessary? Would this not return total number of calories if no meals are found?
		if len(meals) > 0 {

			rows, err = db.Query(context.Background(), "Select SUM(Calories) from meal WHERE dieterid=$1 AND day=$2", dieter.ID, day)
			if err != nil {
				mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
				return 0, errors.New("cannot get today's calories for dieter")
			} else {
				if rows.Next() {
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

// DeleteDieter deletes a dieter from the database.
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

// DeleteMealsForDieter uses a dieter_id to delete all meals for a dieter from the database
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

	// Loop through all meals for the dieter returned from the query above and delete the entries for each meal.
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

		_, err = meal.Query(context.Background(), "DELETE FROM meal WHERE ID=$1", index)
		err = mutils.WrapError(err, "error 101: Cannot delete meal from database", "query")
	}

	return err

}

// GetFoodRow uses a food object to retrieve a food item from the database and return it as a food object
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

// GetMeal uses a meal object to retrieve a meal from the database and return it as a list of meal objects
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

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot parse meal from row returned", err)
		return nil, errors.New("cannot parse meal from row returned")
	}

	return meals, nil
}

// DeleteEntriesByMeal uses a meal_id to delete all entries for a meal from the database
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

// DeleteMeal uses a meal object to find and delete a meal from the database
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

// GetMealCalories uses a meal object to retrieve the total calories for a meal from the database
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

// GetMealEntries uses a meal object to retrieve all entries for a meal from the database
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

// GetDieterMeals uses a dieter object to retrieve all meals for a dieter from the database
func GetDieterMeals(dieter models.Dieter) ([]models.Meal, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * from meal where dieter = $1", dieter.Name)
	err = mutils.WrapError(err, "error 101: Cannot find meals for dieter", "query")
	if err != nil {
		return nil, err
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])
	return meals, mutils.WrapError(err, "error 201: Cannot parse meals from rows", "parse")
}

// AddMeal uses a meal object to add a meal to the database
func AddMeal(meal models.Meal) error {
	var newID int64
	var dieter models.Dieter

	dieter.Name = meal.Dieter
	dieter, err := GetSingleDieter(dieter)
	err = mutils.WrapError(err, "error 101: Cannot find dieter in database", "query")
	if err != nil {
		return err
	}

	meal.Dieterid = dieter.ID

	if meal.Dieterid != 0 {
		db, err := getConnection()
		if err != nil {
			return err
		}

		// Get the meal count to ensure the next ID is unique
		err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from meal").Scan(&newID)
		err = mutils.WrapError(err, "error 102: Cannot get meal count", "query")
		if err != nil {
			return err
		}

		// Add the meal to the database, add 1 to the last ID created
		_, err = db.Exec(context.Background(), "INSERT INTO meal (calories, day, dieter, dieterid, name) values ($1, $2, $3, $4, $5)",
			meal.Calories, meal.Day, meal.Dieter, meal.Dieterid, meal.Name)
		err = mutils.WrapError(err, "error 103: Cannot store new meal", "insert")
		if err != nil {
			return err
		}

		mutils.LogMessage("Request", "Meal added")
		return nil
	}

	return mutils.WrapError(nil, "error 301: Cannot find dieter id", "notfound")
}

// GetAllFood retrieves all food items from the database and returns them as a list of food objects
func GetAllFood() ([]models.Food, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "SELECT * FROM food")
	err = mutils.WrapError(err, "error 101: Cannot query food items", "query")
	if err != nil {
		return nil, err
	}

	food, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Food])
	return food, mutils.WrapError(err, "error 201: Cannot parse food items", "parse")
}

// AddFoodRow uses a food object to add a food item to the database
func AddFoodRow(food models.Food) error {
	db, err := getConnection()
	if err != nil {
		return err
	}

	var count int64
	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from food").Scan(&count)
	err = mutils.WrapError(err, "error 101: Cannot query food count", "query")
	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "INSERT INTO food (id, calories, units, name) values ($1, $2, $3, $4)",
		count+1, food.Calories, food.Units, food.Name)
	return mutils.WrapError(err, "error 102: Cannot insert food", "insert")
}

// UpdateFood uses a food object to update the calories for a food item in the database
func UpdateFood(food models.Food) error {
	db, err := getConnection()
	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "UPDATE food SET Calories = $1 WHERE Name = $2", food.Calories, food.Name)
	return mutils.WrapError(err, "error 101: Cannot update food calories", "update")
}

// DeleteFoodRow uses a food object to find and delete a food item from the database
func DeleteFoodRow(food models.Food) error {
	db, err := getConnection()
	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "DELETE FROM food WHERE Name = $1", food.Name)
	return mutils.WrapError(err, "error 101: Cannot delete food", "delete")
}

// AddEntry uses a complete entry object to add an entry to the database
func AddEntry(entry models.Entry) (models.Entry, error) {
	var newID int64
	db, err := getConnection()
	if err != nil {
		return entry, err
	}

	err = db.QueryRow(context.Background(), "SELECT count(*) AS exact_count from entry").Scan(&newID)
	err = mutils.WrapError(err, "error 101: Cannot query entry count", "query")
	if err != nil {
		return entry, err
	}

	_, err = db.Query(context.Background(), "INSERT INTO entry values ($1, $2, $3, $4)",
		newID+1, entry.Calories, entry.FoodID, entry.MealID)
	err = mutils.WrapError(err, "error 102: Cannot insert entry", "insert")
	if err != nil {
		return entry, err
	}

	return entry, nil
}

// AddEntryToMeal uses an entry object that includes a meal_id to update a meal in the database to add the calories from the entry
func AddEntryToMeal(entry models.Entry) error {
	var meal []models.Meal
	db, err := getConnection()
	if err != nil {
		return err
	}

	meals, err := db.Query(context.Background(), "SELECT * FROM meal WHERE ID = $1", entry.MealID)
	err = mutils.WrapError(err, "error 101: Cannot query meal", "query")
	if err != nil {
		return err
	}

	meal, err = pgx.CollectRows(meals, pgx.RowToStructByName[models.Meal])
	err = mutils.WrapError(err, "error 201: Cannot parse meal", "parse")
	if err != nil {
		return err
	}

	newCalories := entry.Calories + meal[0].Calories
	_, err = db.Query(context.Background(), "UPDATE meal SET Calories = $1 WHERE id = $2", newCalories, entry.MealID)
	return mutils.WrapError(err, "error 102: Cannot update meal calories", "update")
}

// GetEntry uses an entry object to find and return an entry from the database
func GetEntry(entry models.Entry) (models.Entry, error) {
	db, err := getConnection()
	if err != nil {
		return entry, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM entry WHERE ID=$1", entry.ID)
	err = mutils.WrapError(err, "error 101: Cannot query entry", "query")
	if err != nil {
		return entry, err
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Entry])
	err = mutils.WrapError(err, "error 201: Cannot parse entry", "parse")
	if err != nil {
		return entry, err
	}

	return entries[0], nil
}

// DeleteEntry uses an entry object to find and delete an entry from the database
func DeleteEntry(entry models.Entry) error {
	db, err := getConnection()
	if err != nil {
		return err
	}

	_, err = db.Query(context.Background(), "DELETE from ENTRY where ID = $1", entry.ID)
	return mutils.WrapError(err, "error 101: Cannot delete entry", "delete")
}

// retrieveDieter uses a dieter object to find a dieter in the database
func retrieveDieter(dieter models.Dieter) (pgx.Rows, error) {
	db, err := getConnection()
	if err != nil {
		return nil, mutils.WrapError(err, "error 001: Cannot connect to database", "connection")
	}

	rows, err := db.Query(context.Background(), "SELECT * FROM dieter WHERE name = $1", dieter.Name)
	return rows, mutils.WrapError(err, "error 101: Cannot query dieter", "query")
}
