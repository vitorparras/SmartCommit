package main

import (
	"go-microservice/config"
	"go-microservice/routes"
	"log"

	_ "go-microservice/docs" // Importa os documentos gerados pelo Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Backend API
// @version         1.0
// @description     API para gerar mensagens de commit personalizadas usando IA.
// @termsOfService  http://swagger.io/terms/

// @contact.name    Seu Nome
// @contact.url     http://www.seu-site.com
// @contact.email   seu-email@dominio.com

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
// @schemes   http
func main() {
	// Carrega as configurações
	config.LoadConfig()

	// Inicializa o router Gin
	router := gin.Default()

	// Configura as rotas
	routes.SetupRoutes(router)

	// Rota para a documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Executa o servidor
	log.Println("Server running on port", config.ServerPort)
	if err := router.Run(config.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
