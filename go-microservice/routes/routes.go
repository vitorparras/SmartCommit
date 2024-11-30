package routes

import (
	"go-microservice/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheck)
	router.GET("/models", handlers.GetModels)
	router.GET("/commit-patterns", handlers.GetCommitPatterns)
	router.POST("/generate-commit", handlers.GenerateCommit)
}
