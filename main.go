package main

import (
    "context"
	"database/sql"
	"log"
	"mauit/router"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func main() {

    var psqlConn pgx.ConnConfig

    psqlConn.Host = "localhost"
    psqlConn.Port = 5432
    psqlConn.User = "postgres"
    psqlConn.Database = "meal"


    db, err := pgx.Connect(psqlConn)

    if err != nil {
        log.Fatal(err)
    }

    var rows *sql.Rows

    err = db.QueryRow("Select * FROM meal")

    if err != nil {
        log.Fatal(err)
    }

    defer rows.Close()

    for rows.Next() {
        var (
            id int64
            calories int64
        )

        if err := rows.Scan(&id, &calories); err != nil {
            log.Fatal(err)
        }

    }

    r := gin.Default()

    router.SetRoutes(r)

    // start server
    r.Run("localhost:9090")

}
