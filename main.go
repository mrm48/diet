package main

import (
	"context"
	"fmt"
	"log"
	"mauit/models"
	"mauit/router"

	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := pgx.Connect(context.Background(), "postgresql://postgres@localhost:5432/meal")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(context.Background(), "Select * FROM meal")
	meals, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Meal])

	if err != nil {
		log.Fatal(err)
	}

	for _, meal := range meals {
		fmt.Println(meal.ID)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan()
	}

	r := gin.Default()

	router.SetRoutes(r)

	// start server
	r.Run("localhost:9090")

}
