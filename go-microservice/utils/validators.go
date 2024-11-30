package utils

import (
	"errors"
	"go-microservice/models"
	"go-microservice/services/ai"
)

func ValidateCommitRequest(request models.CommitRequest) error {
	// Valida se o modelo é suportado
	if !isValidModel(request.Model) {
		return errors.New("Unsupported AI model")
	}

	// Outras validações conforme necessário
	return nil
}

func isValidModel(model string) bool {
	models := ai.GetAvailableModels()
	for _, m := range models["online"] {
		if m == model {
			return true
		}
	}
	for _, m := range models["offline"] {
		if m == model {
			return true
		}
	}
	return false
}
