package main

import (
	"learning-go-lang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health-check", controllers.HealthCheck)
	r.Run()
}
