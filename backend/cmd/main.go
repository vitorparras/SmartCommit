package main

import (
	_ "backend/docs" // Importa a documentação Swagger
	"backend/internal/api"
	"log"
)

// @title Backend API
// @version 1.0
// @description API para geração de commits e tradução.
// @host localhost:8080
// @BasePath /
func main() {
	router := api.SetupRouter()

	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
