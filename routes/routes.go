package routes

import (
	"fmt"
	"net/http"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println(userID)

	// Fetch user data from database or other source

	c.JSON(http.StatusOK, gin.H{"user": "user_data"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user data to database or other source

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)
}
