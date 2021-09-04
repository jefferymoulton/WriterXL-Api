package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"writerxl-api/models"
)

type CreateProfileInput struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Picture  string `json:"picture"`
}

type ProfileDTO struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}

func LoadProfileRoutes(e *gin.Engine) {
	e.POST("/api/profile", CreateProfile)
	e.GET("/api/profile", FindProfiles)

	e.GET("/api/profile/:email", GetProfile)
}

func CreateProfile(c *gin.Context) {
	var input CreateProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := models.Profile{
		Email:    input.Email,
		Nickname: input.Nickname,
		Name:     input.Name,
		Picture:  input.Picture,
	}
	models.DB.Create(&profile)

	// TODO: Return an error if one occurs (duplicate entry, etc.)
	c.JSON(http.StatusOK, mapDTO(profile))
}

func FindProfiles(c *gin.Context) {
	var profiles []models.Profile
	models.DB.Find(&profiles)

	c.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

func GetProfile(c *gin.Context) {
	var profile models.Profile

	if err := models.DB.Where("email = ?", c.Param("email")).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile was not found"})
		return
	}

	c.JSON(http.StatusOK, mapDTO(profile))
}

func mapDTO(user models.Profile) ProfileDTO {
	var dto ProfileDTO
	dto.Email = user.Email
	dto.Name = user.Name
	dto.Nickname = user.Nickname
	dto.Picture = user.Picture

	return dto
}
