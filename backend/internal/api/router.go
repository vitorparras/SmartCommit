package api

import (
	"log"
	"time"

	"backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Importa os arquivos do Swagger
	ginSwagger "github.com/swaggo/gin-swagger" // Middleware para o Swagger
)

// SetupRouter configura as rotas principais
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware de logs para requisições
	router.Use(func(c *gin.Context) {
		start := time.Now()

		// Processa a requisição
		c.Next()

		// Loga detalhes da requisição e resposta
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		log.Printf(
			"[INFO] %s %s | %d | %s",
			c.Request.Method,
			c.Request.URL.Path,
			statusCode,
			latency,
		)
	})

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rotas principais
	router.POST("/commit", handlers.GenerateCommit)      // Geração de commit
	router.GET("/models", handlers.ListModels)           // Listar modelos disponíveis
	router.GET("/validate", handlers.ValidateRepository) // Validação de repositórios

	return router
}
