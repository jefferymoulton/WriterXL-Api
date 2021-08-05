package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"writerxl-api/routers"
)

func main() {
	r := gin.Default()

	routers.LoadUserRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Printf("Startup service failed, err:%v\n", err)
	}
}
