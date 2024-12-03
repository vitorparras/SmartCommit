package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Model representa um modelo de IA disponível
type Model struct {
	Name        string `json:"name"`
	Type        string `json:"type"` // online ou offline
	Description string `json:"description"`
}

// ListModels retorna uma lista de modelos disponíveis
func ListModels(c *gin.Context) {
	models := []Model{
		{Name: "GPT-2", Type: "offline", Description: "Modelo leve para geração local"},
		{Name: "Bloomz", Type: "offline", Description: "Modelo balanceado para geração local"},
		{Name: "OpenAI GPT-4", Type: "online", Description: "Modelo avançado para geração online"},
	}

	c.JSON(http.StatusOK, models)
}
