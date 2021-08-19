package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"writerxl-api/models"
	"writerxl-api/routes"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"httpw://writerxl.com"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	routes.LoadUserRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Printf("Startup service failed, err:%v\n", err)
	}
}
