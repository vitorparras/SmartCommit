package main

import (
	"log"
	"net/http"

	"backend/internal/adapters/ai"
	"backend/internal/adapters/http/handlers"
	"backend/internal/adapters/repository"
	"backend/internal/core/services"

	_ "backend/docs" // This is required for swagger

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Commit AI API
// @version 1.0
// @description This is a sample server for generating commit messages using AI.
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize adapters
	aiAdapter := ai.NewAIAdapter()
	repoAdapter := repository.NewInMemoryRepository()

	// Initialize service
	commitService := services.NewCommitService(aiAdapter, repoAdapter)

	// Initialize HTTP handler
	handler := handlers.NewHTTPHandler(commitService)

	// Create a new chi router
	r := chi.NewRouter()

	// Use some standard middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Setup CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsMiddleware.Handler)

	// Register our handler
	r.Mount("/", handler)

	// Swagger route
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition
	))

	// Start HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
