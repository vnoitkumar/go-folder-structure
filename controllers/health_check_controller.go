package controllers

import (
	"learning-go-lang/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func HealthCheck(c *gin.Context) {
	healthcheckResponse := responses.HealthCheckResponse{
		Status: "up",
	}
	c.JSON(http.StatusOK, healthcheckResponse)
}
