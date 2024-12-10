package main

import (
	_ "backend/docs" // Importa a documentação Swagger
	"backend/internal/api"
	"log"
	"os/exec"
)

// @title SmartCommit API
// @version 1.0
// @description API para geração de commits, tradução e validação de repositórios.
// @host localhost:8080
// @BasePath /
func main() {
	// Atualiza o Swagger antes de iniciar o servidor
	err := generateSwaggerDocs()
	if err != nil {
		log.Fatalf("Failed to generate Swagger docs: %v", err)
	}

	// Configura o roteador
	router := api.SetupRouter()

	// Inicia o servidor
	log.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// generateSwaggerDocs executa o comando swag init para atualizar a documentação Swagger
func generateSwaggerDocs() error {
	log.Println("Generating Swagger documentation...")

	cmd := exec.Command("swag", "init", "--generalInfo", "cmd/main.go", "--output", "./docs")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Swagger generation failed: %s", string(output))
		return err
	}

	log.Println("Swagger documentation generated successfully.")
	return nil
}
