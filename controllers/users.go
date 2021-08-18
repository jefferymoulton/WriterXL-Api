package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"writerxl-api/models"
)

type CreateUserInput struct {
	ID       uint   `json:"id" binding:"required"`
	AuthID   string `json:"auth_id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
}

func LoadUserRoutes(e *gin.Engine) {
	e.GET("/api/users", FindUsers)

	e.GET("/api/user/:authId", FindUserByAuthId)
	e.PATCH("/api/user/:authId", UpdateUser)

	e.GET("/api/user/email/:email", FindUserByEmail)
}

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUserByAuthId(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("auth_id = ?", c.Param("authId")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User was not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUserByEmail(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("email = ?", c.Param("email")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User was not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("auth_id = ?", c.Param("authId")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User was not found"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
