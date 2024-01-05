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

	cacheRepository := repository.NewCacheUserRepository()
	userCacheService := services.NewCacheUserService(*cacheRepository)
	userCacheRoutes := routes.NewCacheUserRoutes(userCacheService)

	router.GET("/", Welcome)
	router.GET("/ping", Ping)
	router.GET("/users/db/all", userDBRoutes.GetAll)
	router.GET("/users/db/:id", userDBRoutes.GetUser)
	router.POST("/user/db/create", userDBRoutes.CreateUser)
	router.GET("/user/db/drop-user-table", userDBRoutes.DropUserTable)

	router.POST("/user/cache/create", userCacheRoutes.CreateUser)
	router.GET("/user/cache/get-all", userCacheRoutes.GetAll)
	router.GET("/user/cache/get/:cacheID", userCacheRoutes.GetUser)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your REST API!"})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
