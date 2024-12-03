package api

import (
	"backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Importa os arquivos do Swagger
	ginSwagger "github.com/swaggo/gin-swagger" // Middleware para o Swagger
)

// SetupRouter configura as rotas principais
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas principais
	router.POST("/commit", handlers.GenerateCommit) // Geração de commit
	router.GET("/models", handlers.ListModels)      // Listar modelos disponíveis

	return router
}
