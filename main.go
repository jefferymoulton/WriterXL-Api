package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"writerxl-api/models"
	"writerxl-api/routes"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	routes.LoadUserRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Printf("Startup service failed, err:%v\n", err)
	}
}
