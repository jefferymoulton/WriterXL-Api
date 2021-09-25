package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"writerxl-api/models"
)

var path = "/api/profile"

func LoadProfileRoutes(e *gin.Engine) {
	e.POST(path, CreateProfile)

	e.GET(path+"/email/:email", GetProfile)
	e.GET(path+"/:id", GetProfile)

	e.PUT(path+"/:id", UpdateProfile)
	e.PUT(path+"/activate/:id", ActivateProfile)

	e.DELETE(path+"/:id", DeactivateProfile)
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
		profile, err = models.GetProfileByEmail(email)
	}

	if err != nil {
		if err.Error() == "the provided hex string is not a valid ObjectID" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
			return
		} else if err.Error() == "mail: no angle-addr" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address."})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile was not found"})
			return
		}
	}

	c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
	var input models.Profile

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.UpsertProfile(c.Param("id"), input)
	if err != nil && err.Error() == "mongo: no documents in result" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed. Profile was not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func ActivateProfile(c *gin.Context) {
	err := models.ActivateProfile(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}

func DeactivateProfile(c *gin.Context) {
	err := models.DeactivateProfile(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}
