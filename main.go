// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "sonus/m/v2/routes"
)

func main() {
    router := gin.Default()

    // Register routes
    routes.SetupUserRoutes(router)
    // Register other routes here

    // Start the server
    router.Run(":8080")
}
