package handlers

import (
	"net/http"

	"go-microservice/services/ai"

	"github.com/gin-gonic/gin"
)

func GetModels(c *gin.Context) {
	models := ai.GetAvailableModels()
	c.JSON(http.StatusOK, models)
}
