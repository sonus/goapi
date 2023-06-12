// routes/user.go
package routes

import (
    "github.com/gin-gonic/gin"
    "sonus/m/v2/handlers"
)

func SetupUserRoutes(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        userGroup.GET("/:id", handlers.GetUser)
        // Add other user routes here
    }
}
