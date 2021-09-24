package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"writerxl-api/models"
)

var path = "/api/profile"

func LoadProfileRoutes(e *gin.Engine) {
	e.POST(path, CreateProfile)

	e.GET(path+"/:id", GetProfileById)
	e.GET(path+"/email/:email", GetProfileByEmail)

	e.PUT(path+"/:id", UpdateProfile)
}

func CreateProfile(c *gin.Context) {
	var input models.Profile

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateProfile(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "")
}

func GetProfileById(c *gin.Context) {
	var profile models.Profile

	profile, err := models.GetProfileById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile was not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func GetProfileByEmail(c *gin.Context) {
	var profile models.Profile

	profile, err := models.GetProfileByEmail(c.Param("email"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile was not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
	var input models.Profile

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.UpdateProfile(c.Param("id"), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
