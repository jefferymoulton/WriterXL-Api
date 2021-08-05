package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func userHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"firstName": "Jeff",
		"lastName":  "Moulton"})
}

func LoadUserRoutes(e *gin.Engine) {
	e.GET("/api/user", userHandler)
}
