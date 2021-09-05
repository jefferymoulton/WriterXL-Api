package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"writerxl-api/dto"
	"writerxl-api/models"
)

type ProfileDTO struct {
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

var path = "/api/profile"

func LoadProfileRoutes(e *gin.Engine) {
	e.POST(path, CreateProfile)
	e.PUT(path, UpdateProfile)

	e.GET(path+"/:email", GetProfile)
}

func CreateProfile(c *gin.Context) {
	var input models.ProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.CreateProfile(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapDTO(profile))
}

func UpdateProfile(c *gin.Context) {
	var input models.ProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.UpdateProfile(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapDTO(profile))
}

func GetProfile(c *gin.Context) {
	var profile models.Profile

	profile, err := models.GetProfileByEmail(c.Param("email"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile was not found"})
		return
	}

	c.JSON(http.StatusOK, dto.MapDTO(profile))
}
