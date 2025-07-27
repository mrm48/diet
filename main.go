package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"mauit/mutils"
	"mauit/router"
	"net/http"
	"os"
	"time"
	"strconv"
)

// main initializes and starts mauit. Key responsibilities:
// - Creates/opens log file at logs/mauit_app.log with append mode
// - Sets up all API routes via router.SetRoutes()
// - Starts HTTP server on localhost:9090
func main() {

	host := "localhost"
	portNumber := 9090
	routerHost := host + strconv.Itoa(portNumber)

	// create and / or open the application log file.
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
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	mutils.LogMessage("Server Startup", "Initializing")

	// setup the router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// Setup API routes
	router.SetRoutes(r)

	// Serve static frontend files
	r.StaticFS("/static", http.Dir("./frontend"))

	// Serve index.html at the root path
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	mutils.LogMessage("Server Startup", "Routes set and frontend configured: Starting server")

	// start server
	r.Run(routerHost)

}
