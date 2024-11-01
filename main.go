package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mauit/router"
	"os"
)

func main() {
	f, err := os.OpenFile("mauit_app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening log file: %v\n", err)
		return
	}
	defer f.Close()
	log.SetOutput(f)
	log.Output(1, "Initializing")
	r := gin.Default()

	router.SetRoutes(r)

	// start server
	r.Run("localhost:9090")

}
