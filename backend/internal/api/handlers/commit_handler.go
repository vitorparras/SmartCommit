package handlers

import (
	"backend/pkg/commit"
	"backend/pkg/translation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateCommit processa a geração de uma mensagem de commit
func GenerateCommit(c *gin.Context) {
	var req struct {
		Pattern  string `json:"pattern"`
		Message  string `json:"message"`
		Model    string `json:"model"`
		Language string `json:"language"`
	}

	// Bind JSON para a requisição
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Geração de mensagem de commit
	commitMessage, err := commit.Generate(req.Pattern, req.Message, req.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate commit message"})
		return
	}

	// Tradução (se necessário)
	if req.Language != "" {
		translatedMessage, err := translation.Translate(commitMessage, req.Language)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate commit message"})
			return
		}
		commitMessage = translatedMessage
	}

	c.JSON(http.StatusOK, gin.H{"commitMessage": commitMessage})
}
