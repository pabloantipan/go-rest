package routes

import (
	"fmt"
	"net/http"

	"example.com/mod/models"
	"example.com/mod/repository"
	"example.com/mod/services"
	"github.com/gin-gonic/gin"
)

type UserRoutes interface {
	GetAll(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userRoutesImpl struct {
	userService services.UserService
}

func NewUserRoutes(userService services.UserService) UserRoutes {
	return &userRoutesImpl{userService: userService}
}

func (u *userRoutesImpl) GetAll(c *gin.Context) {
	users, err := u.userService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *userRoutesImpl) GetUser(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println(userID)

	// Fetch user data from database or other source
	repository.Connect()

	c.JSON(http.StatusOK, gin.H{"user": "user_data"})
}

func (u *userRoutesImpl) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user data to database or other source

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println(userID)

	// Fetch user data from database or other source
	repository.Connect()

	c.JSON(http.StatusOK, gin.H{"user": "user_data"})
}
