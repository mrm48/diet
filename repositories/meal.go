package repositories

import (
	"context"

	"mauit/mutils"
    "mauit/models"

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
