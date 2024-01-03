package main

import (
	"net/http"

	"example.com/mod/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to your REST API!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users/:id", routes.GetUser)

	router.POST("/user/create", routes.CreateUser)

	router.Run() // listen and serve on 0.0.0.0:8080
}
