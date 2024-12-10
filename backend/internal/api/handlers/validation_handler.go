package handlers

import (
	"backend/pkg/validation"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidationError representa o formato de erro de validação
type ValidationError struct {
	Error string `json:"error"`
}

// @Summary Valida um repositório local
// @Description Verifica se o caminho fornecido é válido e contém um repositório Git
// @Tags Validation
// @Accept json
// @Produce json
// @Param path query string true "Caminho do repositório"
// @Success 200 {object} validation.ValidationResult
// @Failure 400 {object} ValidationError
// @Failure 500 {object} ValidationError
// @Router /validate [get]
func ValidateRepository(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, ValidationError{Error: "Path is required"})
		log.Printf("[ERROR] Validation failed: Path is required")
		return
	}

	result, err := validation.ValidateRepository(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ValidationError{Error: "Internal server error"})
		log.Printf("[ERROR] Validation failed: %v", err)
		return
	}

	c.JSON(http.StatusOK, result)
	log.Printf("[INFO] Validation successful for path: %s", path)
}
