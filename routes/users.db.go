package routes

import (
	"fmt"
	"net/http"

	"example.com/mod/models"
	"example.com/mod/services"
	"github.com/gin-gonic/gin"
)

type UserDBRoutes interface {
	GetAll(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	DropUserTable(c *gin.Context)
}

type userDBRoutesImpl struct {
	userDBService services.UserDBService
}

func NewDBUserRoutes(userDBService services.UserDBService) UserDBRoutes {
	return &userDBRoutesImpl{userDBService: userDBService}
}

func (u *userDBRoutesImpl) GetAll(c *gin.Context) {
	users, err := u.userDBService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *userDBRoutesImpl) GetUser(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println(userID)

	// Fetch user data from database or other source

	c.JSON(http.StatusOK, gin.H{"user": "user_data"})
}

func (u *userDBRoutesImpl) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"parse error": err.Error()})
		return
	}

	result, err := u.userDBService.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"repo error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully!",
		"result":  result,
	})
}

func (u *userDBRoutesImpl) DropUserTable(c *gin.Context) {
	u.userDBService.DropUserTable()
}
