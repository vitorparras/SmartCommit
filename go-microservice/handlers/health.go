// handlers/health.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck
// @Summary      Verifica o status do servidor
// @Description  Retorna o status de saúde do servidor
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Server is running",
	})
}
