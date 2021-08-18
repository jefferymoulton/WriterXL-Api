package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"writerxl-api/controllers"
	"writerxl-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	controllers.LoadUserRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Printf("Startup service failed, err:%v\n", err)
	}
}
