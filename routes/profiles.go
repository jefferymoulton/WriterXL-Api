package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	"writerxl-api/models"
)

var path = "/api/profile"

func LoadProfileRoutes(e *gin.Engine) {
	e.POST(path, CreateProfile)

	e.GET(path+"/:email", GetProfile)
	e.GET(path+"/id/:id", GetProfile)

	e.PUT(path+"/:email", UpdateProfile)
}

func CreateProfile(c *gin.Context) {
	var input models.Profile

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.CreateProfile(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

func GetProfile(c *gin.Context) {
	var profile models.Profile
	var err error

	email := c.Param("email")
	id := c.Param("id")

	if id != "" {
		profile, err = models.GetProfileById(id)
	} else {
		_, err := mail.ParseAddress(email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address."})
			return
		}

		profile, err = models.GetProfileByEmail(email)
	}
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

	profile, err := models.UpsertProfile(c.Param("email"), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
