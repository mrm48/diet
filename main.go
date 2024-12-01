package main

import (
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
	"log"
	"mauit/mutils"
	"mauit/router"
	"os"
)

func main() {

	f, err := os.OpenFile("logs/mauit_app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(f)

	log.SetOutput(f)
	mutils.LogMessage("Server Startup", "Initializing")
	r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET","PUT","POST","DELETE","PATCH"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:5173"
        },
        MaxAge: 12 * time.Hour,
    }))

	router.SetRoutes(r)
	mutils.LogMessage("Server Startup", "Routes set: Starting server")
	// start server
	r.Run("localhost:9090")

}

