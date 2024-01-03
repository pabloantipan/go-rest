package main

import (
	"net/http"

	"example.com/mod/repository"
	"example.com/mod/routes"
	"example.com/mod/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := repository.Connect()

	repository := repository.NewUserRepository(db)

	userService := services.NewUserService(*repository)

	userRoutes := routes.NewUserRoutes(userService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to your REST API!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users/:id", userRoutes.GetUser)

	router.POST("/user/create", userRoutes.CreateUser)

	router.Run() // listen and serve on 0.0.0.0:8080
}
