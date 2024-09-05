package main

import (
	"mauit/router"

	"github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()

    router.SetRoutes(r)

    // start server
    r.Run("localhost:9090")

}
