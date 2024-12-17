package ai

import (
	"backend/internal/core/domain"
	"fmt"
)

type AIAdapter struct {
	// Configuração do serviço de IA iria aqui
}

func NewAIAdapter() *AIAdapter {
	return &AIAdapter{}
}

func (a *AIAdapter) GenerateCommitMessage(description string, readmeContent string, config domain.CommitConfig) (string, error) {
	// Aqui você implementaria a chamada real para o seu serviço de IA
	// usando os parâmetros de configuração para personalizar a geração
	return fmt.Sprintf("[%s] AI generated commit for: %s (Style: %s, Lang: %s, Readme length: %d)",
		config.Format,
		description,
		config.Style,
		config.Language,
		len(readmeContent),
	), nil
}

func (a *AIAdapter) GetAvailableModels() []string {
	return []string{
		"gpt-3.5-turbo",
		"gpt-4",
		"codex",
		"claude-2",
	}
}
