package commit

import (
	"errors"
	"fmt"
)

// Generate gera uma mensagem de commit baseada em um padrão, mensagem e modelo
func Generate(pattern, message, model string) (string, error) {
	// Recupera o padrão correspondente
	template, exists := GetTemplate(pattern)
	if !exists {
		return "", errors.New("invalid commit pattern")
	}

	// Substitui o placeholder no padrão pela mensagem
	commitMessage := fmt.Sprintf(template, map[string]string{
		"description": message,
	})

	// Simulação de integração com modelo (online/offline)
	commitMessage = applyModelLogic(commitMessage, model)

	return commitMessage, nil
}

// applyModelLogic simula o uso de IA (online/offline) para refinar a mensagem
func applyModelLogic(message, model string) string {
	// Lógica simulada: Adiciona o nome do modelo ao final da mensagem
	return fmt.Sprintf("%s (Generated by model: %s)", message, model)
}