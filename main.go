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

	dbRepository := repository.NewDBUserRepository()

	userDBService := services.NewDBUserService(*dbRepository)

	userDBRoutes := routes.NewDBUserRoutes(userDBService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to your REST API!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users/db/all", userDBRoutes.GetAll)

	router.GET("/users/db/:id", userDBRoutes.GetUser)

	router.POST("/user/db/create", userDBRoutes.CreateUser)

	router.GET("/user/db/drop-user-table", userDBRoutes.DropUserTable)

	router.Run() // listen and serve on 0.0.0.0:8080
}
