package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"writerxl-api/models"
)

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Picture  string `json:"picture"`
}

type UserDTO struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}

func LoadUserRoutes(e *gin.Engine) {
	e.POST("/api/user", CreateUser)
	e.GET("/api/user", FindUsers)

	e.GET("/api/user/email/:email", GetUserByEmail)
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Nickname: input.Nickname,
		Name:     input.Name,
		Picture:  input.Picture,
	}
	models.DB.Create(&user)

	// TODO: Return an error if one occurs (duplicate entry, etc.)
	c.JSON(http.StatusOK, mapDTO(user))
}

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByEmail(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("email = ?", c.Param("email")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User was not found"})
		return
	}

	c.JSON(http.StatusOK, mapDTO(user))
}

func mapDTO(user models.User) UserDTO {
	var dto UserDTO
	dto.Email = user.Email
	dto.Name = user.Name
	dto.Nickname = user.Nickname
	dto.Picture = user.Picture

	return dto
}

/*
func UpdateUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("auth_id = ?", c.Param("authId")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User was not found"})
		return
	}

	var input UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"user": user})
}*/
