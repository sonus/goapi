// handlers/user.go
package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sonus/m/v2/models"
)

func GetUser(c *gin.Context) {
    // Logic to fetch user from database
    user := models.User{
        ID:    1,
        Name:  "John Doe",
        Email: "john@example.com",
    }

    c.JSON(http.StatusOK, user)
}

// Add other user-related handlers here
