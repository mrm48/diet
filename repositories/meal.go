package repositories

import (
	"context"

	"mauit/mutils"
    "mauit/models"

    "strconv"

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

func GetAllDieters() ([]models.Dieter, error) {
    
    db, err := getConnection()
	rows, err := db.Query(context.Background(), "Select * FROM dieter")

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter rows from database", err)
		return nil, err
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create list of dieters from rows returned", err)
		return nil, err
	}

	defer rows.Close()
    return Dieters, err

}

func GetSingleDieter(dieter models.Dieter) (models.Dieter, error) {

	db, err := getConnection()

	if err != nil {
		return dieter, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE name=$1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get dieter information", err)
		return dieter, err
	}

	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

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

func AddNewDieter(dieter models.Dieter) (error) {

	var newID int64

	db, err := getConnection()

	if err != nil {
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

func UpdateDieterCalories(dieter models.Dieter) (error) {

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
		return err
	}

    return nil
}

func GetDieterCalories (dieter models.Dieter) ([]models.Dieter, error) {
	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM dieter WHERE Name = $1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
		return nil, err
	}
	Dieters, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot create a dieter object from search", err)
		return nil, err
	}

    return Dieters, nil
}

func GetDieterMealsToday(dieter models.Dieter, day string) ([]models.Meal, error) {
	db, err := getConnection()

	if err != nil {
		return nil, err
	}

    rows, err := db.Query(context.Background(), "SELECT * from meal WHERE dieter=$1 AND day=$2", dieter.Name, day)
    
    if err != nil {
        mutils.LogApplicationError("Database Error", "Cannot retrieve meals by day for dieter from database", err)
        return nil, err
    }

    meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

    if err != nil {
        mutils.LogApplicationError("Application Error", "Cannot populate list of meals with data returned from database", err)
        return nil, err
    }

    return meals, nil
}

func GetRemainingCaloriesToday(dieter models.Dieter, day string) (int, error) {
	db, err := getConnection()

	if err != nil {
		return 0, err
	}

	rows, err := db.Query(context.Background(), "Select * from dieter WHERE Name = $1", dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Database error", "Cannot find dieter by name", err)
		return 0, err
	}

	Dieter, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dieter])

	if Dieter != nil {

        rows, err := db.Query(context.Background(), "SELECT * from meal WHERE dieterid=$1 AND day=$2,", dieter.ID, day)
        
	    meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

        if len(meals) > 0 {

    		rows, err = db.Query(context.Background(), "Select SUM(Calories) from meal WHERE dieterid=$1 AND day=$2", dieter.ID, day)
    		if err != nil {
    			mutils.LogApplicationError("Database Error", "Cannot retrieve dieter information from database", err)
    			return 0, err
    		} else {
    			if rows.Next() == true {
    				err = rows.Scan(&dieter.Calories)
    				if err != nil {
    					mutils.LogApplicationError("Request", "Cannot parse sum of calories for this dieter", err)
    					return 0, err
    				} else {
    					return Dieter[0].Calories-dieter.Calories, err
    				}
    			}
    		}
        } else {
            return Dieter[0].Calories, nil
        }
	} else {
		mutils.LogApplicationError("Database Error", "Cannot find remaining dieter calories requested", nil)
		return 0, err
	}

    return 0, err
}

func DeleteDieter(dieter models.Dieter) error {
	db, err := getConnection()

    if err != nil {
        return err
    }

	err = db.QueryRow(context.Background(), "SELECT * from dieter WHERE Name=$1", dieter.Name).Scan(&dieter.ID, &dieter.Calories, &dieter.Name)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot retrieve dieter with name provided", err)
		return err
	}

	DeleteMealsForDieter(dieter.ID)

	_, err = db.Query(context.Background(), "DELETE from dieter WHERE ID=$1", dieter.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot delete dieter retrieved by ID", err)
		return err
	}

    return nil
}

func DeleteMealsForDieter(dieterID int64) error {

	meal, err := getConnection()

	if err != nil {
		return err
	}

	rows, err := meal.Query(context.Background(), "SELECT ID FROM meal WHERE dieterID=$1", dieterID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot find meals by dieter from database", err)
		return err
	}

	defer rows.Close()

	var index int64

	for rows.Next() {
		err = rows.Scan(&index)
		if err != nil {
			mutils.LogApplicationError("Application Error", "Cannot get meal ID from returned rows", err)
			return err
		}
        err = DeleteEntriesByMeal(index)

        if err != nil {
            mutils.LogApplicationError("Application Error", "Cannot delete entries for meal", err)
            return err
        }

		conn, err := getConnection()
		if err != nil {
			return err
		}
		_, err = conn.Query(context.Background(), "DELETE FROM meal WHERE ID=$1", index)
		if err != nil {
			mutils.LogApplicationError("Database Error", "Cannot delete meal from database", err)
			return err
		}
	}

	return nil

}

func GetFoodRow(food models.Food) (models.Food, error) {

    var errorFood models.Food
    errorFood.Name = "nil"

	db, err := getConnection()

	rows, err := db.Query(context.Background(), "Select * FROM Food WHERE name=$1", food.Name)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot get food information", err)
		return errorFood, err
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Food])

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

func GetMeal (meal models.Meal) ([]models.Meal, error) {
	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(context.Background(), "Select * FROM meal WHERE name=$1 AND dieter=$2 AND day=$3", meal.Name, meal.Dieter, meal.Day)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot query meal from database", err)
		return nil, err
	}

	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

    return meals, nil
}


func DeleteEntriesByMeal(mealID int64) error {

	meal, err := getConnection()

	if err != nil {
		return err
	}

	_, err = meal.Query(context.Background(), "DELETE FROM entry WHERE MEAL_ID=$1", mealID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot delete entries for meal from database", err)
		return err
	}

    return nil

}

func DeleteMeal(meal models.Meal) error {
	db, err := getConnection() 

	if err != nil {
		return err
	}

	var dbMeal models.Meal
	err = db.QueryRow(context.Background(), "SELECT ID FROM meal WHERE Name = $1 AND Dieter = $2 AND Day = $3", meal.Name, meal.Dieter, meal.Day).Scan(&dbMeal.ID)

	if err != nil {
		mutils.LogApplicationError("Application Error", "Cannot find meal in database", err)
		return err
	}

	meal.ID = dbMeal.ID

	DeleteEntriesByMeal(meal.ID)

	_, err = db.Query(context.Background(), "DELETE FROM meal WHERE ID=$1", meal.ID)

	if err != nil {
		mutils.LogApplicationError("Database Error", "Cannot meal food from database", err)
		return err
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

func AddMeal(meal models.Meal) (error) {
	var newID int64
    var dieter models.Dieter

	// add in any missing fields to meal object (don't need day, dieterid or calories)
	if meal.Day == "" {
		meal.Day = models.GetCurrentDate()
	}

    dieter.Name = meal.Dieter

    GetSingleDieter(dieter)

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

    db, err  := getConnection()

	if err != nil {
        return err
	}

	_, err = db.Query(context.Background(), "DELETE FROM food WHERE Name = $1", food.Name)

    return err
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
