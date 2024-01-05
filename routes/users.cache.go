package routes

import (
	"net/http"

	"example.com/mod/models"
	"example.com/mod/services"
	"github.com/gin-gonic/gin"
)

type UserCacheRoutes interface {
	GetAll(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userCacheRoutesImpl struct {
	userCacheService services.UserCacheService
}

func NewCacheUserRoutes(userCacheService services.UserCacheService) UserCacheRoutes {
	return &userCacheRoutesImpl{userCacheService: userCacheService}
}

func (u *userCacheRoutesImpl) GetAll(c *gin.Context) {
	users, err := u.userCacheService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *userCacheRoutesImpl) GetUser(c *gin.Context) {
	userID := c.Param("cacheID")

	user, err := u.userCacheService.GetByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *userCacheRoutesImpl) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"parse error": err.Error()})
		return
	}

	err := u.userCacheService.Create(user)

	c.JSON(http.StatusOK, gin.H{"result": err})
	return
}
